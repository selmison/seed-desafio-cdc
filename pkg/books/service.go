package books

import (
	"context"
	"fmt"
	"log"
	"strings"

	kitLog "github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/books"
)

type service struct {
	repo   *gorm.DB
	logger kitLog.Logger
}

// NewService returns a new Books Service.
func NewService(repo *gorm.DB, logger kitLog.Logger) books.Service {
	return &service{repo, logger}
}

// CreateBook implements create.
func (s *service) CreateBook(_ context.Context, dto *books.CreateBookDTO) (res *books.BookDTO, err error) {
	summary := ""
	if dto.Summary != nil {
		summary = strings.TrimSpace(*dto.Summary)
	}
	book := Book{
		ID:         uuid.New().String(),
		Title:      strings.TrimSpace(dto.Title),
		Synopsis:   strings.TrimSpace(dto.Synopsis),
		Summary:    summary,
		Price:      dto.Price,
		Pages:      dto.Pages,
		Isbn:       strings.TrimSpace(dto.Isbn),
		Issue:      strings.TrimSpace(dto.Issue),
		CategoryID: strings.TrimSpace(dto.CategoryID),
		ActorID:    strings.TrimSpace(dto.ActorID),
	}
	if err := book.Validate(); err != nil {

		return nil, err
	}
	if err := s.logger.Log("info", fmt.Sprintf("books.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	result := s.repo.Create(&book)
	return &books.BookDTO{
		ID:         book.ID,
		Title:      book.Title,
		Synopsis:   book.Synopsis,
		Summary:    &summary,
		Price:      dto.Price,
		Pages:      dto.Pages,
		Isbn:       book.Isbn,
		Issue:      book.Issue,
		CategoryID: book.CategoryID,
		ActorID:    book.ActorID,
	}, result.Error
}
