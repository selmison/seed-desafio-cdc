package catalog

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// CreateBook implements create_book.
func (s *service) CreateBook(_ context.Context, dto *catalogGen.CreateBookDTO) (res *catalogGen.BookDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("books.create_book")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	book := s.mapCreateBookDTOToBook(*dto)
	if err := book.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&book); result.Error != nil {
		return nil, result.Error
	}
	actorIDs := s.mapActorsToIDs(book.Actors)
	categoryIDs := s.mapCategoriesToIDs(book.Categories)
	return &catalogGen.BookDTO{
		ID:          book.ID,
		Title:       book.Title,
		Synopsis:    book.Synopsis,
		Summary:     &book.Summary,
		Price:       book.Price,
		Pages:       book.Pages,
		Isbn:        book.Isbn,
		Issue:       book.Issue,
		CategoryIds: categoryIDs,
		ActorIds:    actorIDs,
	}, nil
}

// ListBook implements list_book.
func (s *service) ListBooks(context.Context) (res []*catalogGen.BookDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("books.list_books")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	var books []*Book
	result := s.repo.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	bookDTOS := make([]*catalogGen.BookDTO, len(books))
	for i, book := range books {
		dto, err := s.mapBookToBookDTO(*book)
		if err != nil {
			return nil, err
		}
		bookDTOS[i] = dto
	}
	return bookDTOS, nil
}

// ShowBook implements show_book.
func (s *service) ShowBook(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.BookDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalogGen.show_book")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	book := Book{}
	result := s.repo.Where("id = ?", dto.ID).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	return s.mapBookToBookDTO(book)
}
