package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	goa "goa.design/goa/v3/pkg"
	"gorm.io/gorm"

	actorsGen "github.com/selmison/seed-desafio-cdc/gen/actors"
	booksGen "github.com/selmison/seed-desafio-cdc/gen/books"
	categoriesGen "github.com/selmison/seed-desafio-cdc/gen/categories"
	"github.com/selmison/seed-desafio-cdc/pkg/actors"
	"github.com/selmison/seed-desafio-cdc/pkg/books"
	"github.com/selmison/seed-desafio-cdc/pkg/categories"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

func ValidationCreateActorMiddleware(repo *gorm.DB) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			dto := request.(*actorsGen.CreateActorDTO)
			email := strings.TrimSpace(dto.EMail)
			if err := fieldShouldBeUnique(repo, "e-mail", email, actors.Actor{}); err != nil {
				return nil, err
			}
			length := 400
			if len(dto.Description) > length {
				return nil, goa.InvalidLengthError("body.description", dto.Description, len(dto.Description), length, false)
			}
			return next(ctx, request)
		}
	}
}

func ValidationCreateCategoryMiddleware(repo *gorm.DB) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			dto := request.(*categoriesGen.CreateCategoryDTO)
			name := strings.TrimSpace(dto.Name)
			if err := fieldShouldBeUnique(repo, "name", name, categories.Category{}); err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}

func ValidationBookMiddleware(repo *gorm.DB) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if dto, ok := request.(*booksGen.CreateBookDTO); ok {
				if err := fieldShouldBeUnique(repo, "title", strings.TrimSpace(dto.Title), books.Book{}); err != nil {
					return nil, err
				}
				if err := fieldShouldBeUnique(repo, "isbn", strings.TrimSpace(dto.Isbn), books.Book{}); err != nil {
					return nil, err
				}
				length := 500
				if len(dto.Synopsis) > length {
					return nil, goa.InvalidLengthError("body.synopsis", dto.Synopsis, len(dto.Synopsis), length, false)
				}
				issue, err := time.Parse(time.RFC3339, dto.Issue)
				if err != nil {
					return nil, err
				}
				today := time.Now()
				if today.After(issue) {
					return nil, &goa.ServiceError{
						Name:    "invalid_fields",
						ID:      goa.NewErrorID(),
						Message: fmt.Sprintf("the '%s' %s", "body.issue", "should be in the future"),
					}
				}
			}
			return next(ctx, request)
		}
	}
}

func fieldShouldBeUnique(repo *gorm.DB, fieldName, fieldValue string, result interface{}) error {
	query := fmt.Sprintf("%s = ?", fieldName)
	var rows int64
	switch v := result.(type) {
	case actors.Actor:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	case categories.Category:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	case books.Book:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	default:
		return fmt.Errorf("uniqueness validation: type %w", coreDomain.ErrIsNotValid)
	}
	if rows > 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.%s' %v", fieldName, coreDomain.ErrShouldBeUnique),
		}
	}
	return nil
}
