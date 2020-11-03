package books

import (
	"context"
	"fmt"
	"log"
	"strings"

	kitLog "github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	booksGen "github.com/selmison/seed-desafio-cdc/gen/books"
)

type service struct {
	repo   *gorm.DB
	logger kitLog.Logger
}

// NewService returns a new Books Service.
func NewService(repo *gorm.DB, logger kitLog.Logger) booksGen.Service {
	return &service{repo, logger}
}

// CreateBook implements create.
func (s *service) CreateBook(_ context.Context, dto *booksGen.CreateBookDTO) (res *booksGen.BookDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("books.create_book")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
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
	result := s.repo.Create(&book)
	return &booksGen.BookDTO{
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

// ListBook implements create.
func (s *service) ListBooks(context.Context) (res []*booksGen.BookDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("books.list_books")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	var books []*Book
	result := s.repo.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	bookDTOS := make([]*booksGen.BookDTO, len(books))
	for i, book := range books {
		bookDTOS[i] = &booksGen.BookDTO{
			ID:         book.ID,
			Title:      book.Title,
			Synopsis:   book.Synopsis,
			Summary:    &book.Summary,
			Price:      book.Price,
			Pages:      book.Pages,
			Isbn:       book.Isbn,
			Issue:      book.Issue,
			CategoryID: book.CategoryID,
			ActorID:    book.ActorID,
		}
	}
	return bookDTOS, nil
}
