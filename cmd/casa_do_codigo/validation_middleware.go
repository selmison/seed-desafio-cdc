package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	goa "goa.design/goa/v3/pkg"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	"github.com/selmison/seed-desafio-cdc/pkg/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

func ValidationMiddleware(repo *gorm.DB) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			switch dto := request.(type) {
			case *catalogGen.CreateActorDTO:
				if err := CreateActorValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreateBookDTO:
				if err := CreateBookValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreateCategoryDTO:
				if err := CreateCategoryValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreateCouponDTO:
				if err := CreateCouponValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreateCustomerDTO:
				if err := CreateCustomerValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreatePurchaseDTO:
				if err := CreatePurchaseValidation(repo, dto); err != nil {
					return nil, err
				}
			case *catalogGen.CreateStateDTO:
				if err := CreateStateValidation(repo, dto); err != nil {
					return nil, err
				}
			}
			return next(ctx, request)
		}
	}
}

func CreateActorValidation(repo *gorm.DB, dto *catalogGen.CreateActorDTO) error {
	email := strings.TrimSpace(dto.Email)
	if err := fieldShouldBeUnique(repo, "email", email, catalog.Actor{}); err != nil {
		return err
	}
	length := 400
	if len(dto.Description) > length {
		return goa.InvalidLengthError("body.description", dto.Description, len(dto.Description), length, false)
	}
	return nil
}

func CreateCategoryValidation(repo *gorm.DB, dto *catalogGen.CreateCategoryDTO) error {
	name := strings.TrimSpace(dto.Name)
	if err := fieldShouldBeUnique(repo, "name", name, catalog.Category{}); err != nil {
		return err
	}
	return nil
}

func CreateBookValidation(repo *gorm.DB, dto *catalogGen.CreateBookDTO) error {
	if err := fieldShouldBeUnique(repo, "title", strings.TrimSpace(dto.Title), catalog.Book{}); err != nil {
		return err
	}
	if err := fieldShouldBeUnique(repo, "isbn", strings.TrimSpace(dto.Isbn), catalog.Book{}); err != nil {
		return err
	}
	length := 500
	if len(dto.Synopsis) > length {
		return goa.InvalidLengthError("body.synopsis", dto.Synopsis, len(dto.Synopsis), length, false)
	}
	issue, err := time.Parse(time.RFC3339, dto.Issue)
	if err != nil {
		return err
	}
	today := time.Now()
	if today.After(issue) {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the '%s' %s", "body.issue", "should be in the future"),
		}
	}
	return nil
}

func CreateCouponValidation(repo *gorm.DB, dto *catalogGen.CreateCouponDTO) error {
	if err := fieldShouldBeUnique(repo, "code", strings.TrimSpace(dto.Code), catalog.Coupon{}); err != nil {
		return err
	}
	if dto.Discount <= 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the '%s' %s", "body.discount", "should be in the future"),
		}
	}
	issue, err := time.Parse(time.RFC3339, dto.Validity)
	if err != nil {
		return err
	}
	today := time.Now()
	if today.After(issue) {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.validity' value %v", coreDomain.ErrIsNotValid),
		}
	}
	return nil
}

func CreateCustomerValidation(repo *gorm.DB, dto *catalogGen.CreateCustomerDTO) error {
	result := repo.Where("id = ?", dto.Address.StateID).First(&catalog.State{})
	if result.RowsAffected == 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.customer.address.state_id' was %v", coreDomain.ErrNotFound),
		}
	}
	return nil
}

func CreatePurchaseValidation(repo *gorm.DB, dto *catalogGen.CreatePurchaseDTO) error {
	if err := CreateCustomerValidation(repo, dto.Customer); err != nil {
		return err
	}
	if dto.Cart.CouponID != nil && strings.TrimSpace(*dto.Cart.CouponID) != "" {
		result := repo.Where("id = ?", dto.Cart.CouponID).First(&catalog.Coupon{})
		if result.RowsAffected == 0 {
			return &goa.ServiceError{
				Name:    "invalid_fields",
				ID:      goa.NewErrorID(),
				Message: fmt.Sprintf("the 'body.cart.coupon_id' was %v", coreDomain.ErrNotFound),
			}
		}
	}
	if dto.Cart.Total <= 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.cart.total' was %v", coreDomain.ErrIsNotValid),
		}
	}
	for i, item := range dto.Cart.Items {
		result := repo.Where("id = ?", item.BookID).First(&catalog.Book{})
		if result.RowsAffected == 0 {
			return &goa.ServiceError{
				Name:    "invalid_fields",
				ID:      goa.NewErrorID(),
				Message: fmt.Sprintf("the 'body.cart.items[%d].book_id' was %v", i, coreDomain.ErrNotFound),
			}
		}
	}
	return TotalCartValidation(repo, dto.Cart)
}

func CreateStateValidation(repo *gorm.DB, dto *catalogGen.CreateStateDTO) error {
	result := repo.Where("id = ?", dto.CountryID).First(&catalog.Country{})
	if result.RowsAffected == 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.country_id' was %v", coreDomain.ErrNotFound),
		}
	}
	return nil
}

func TotalCartValidation(repo *gorm.DB, dto *catalogGen.CreateCartDTO) error {
	totals := make([]float32, len(dto.Items))
	for i, item := range dto.Items {
		book := &catalog.Book{}
		result := repo.Where("id = ?", item.BookID).First(book)
		if result.RowsAffected == 0 {
			return &goa.ServiceError{
				Name:    "invalid_fields",
				ID:      goa.NewErrorID(),
				Message: fmt.Sprintf("the 'body.cart.item[%d].book_id' was %v", i, coreDomain.ErrNotFound),
			}
		}
		totals[i] = float32(item.Amount) * book.Price
	}
	var total float32 = 0
	for _, price := range totals {
		total = total + price
	}
	if total != dto.Total {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.cart.total' (%f) should be equal the calculated total (%f)", dto.Total, total),
		}
	}
	return nil
}

func fieldShouldBeUnique(repo *gorm.DB, fieldName, fieldValue string, result interface{}) error {
	query := fmt.Sprintf("%s = ?", fieldName)
	var rows int64
	switch v := result.(type) {
	case catalog.Actor:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	case catalog.Book:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	case catalog.Category:
		rows = repo.Where(query, fieldValue).First(&v).RowsAffected
	case catalog.Coupon:
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
