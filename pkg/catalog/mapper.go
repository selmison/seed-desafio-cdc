package catalog

import (
	"strings"

	"github.com/google/uuid"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

func (s *service) mapCreateBookDTOToBook(dto catalogGen.CreateBookDTO) Book {
	var summary string
	if dto.Summary != nil {
		summary = strings.TrimSpace(*dto.Summary)
	}
	actors := s.mapIDsToActors(dto.ActorIds)
	categories := s.mapIDsToCategories(dto.CategoryIds)
	book := Book{
		ID:         uuid.New().String(),
		Title:      strings.TrimSpace(dto.Title),
		Synopsis:   strings.TrimSpace(dto.Synopsis),
		Summary:    summary,
		Price:      dto.Price,
		Pages:      dto.Pages,
		Isbn:       strings.TrimSpace(dto.Isbn),
		Issue:      strings.TrimSpace(dto.Issue),
		Categories: categories,
		Actors:     actors,
	}
	return book
}

func (s *service) mapBookToBookDTO(book Book) (*catalogGen.BookDTO, error) {
	var actors []*Actor
	if err := s.repo.Model(&book).Association("Actors").Find(&actors); err != nil {
		return nil, err
	}
	var categories []*Category
	if err := s.repo.Model(&book).Association("Categories").Find(&categories); err != nil {
		return nil, err
	}
	actorIDs := s.mapActorsToIDs(actors)
	categoryIDs := s.mapCategoriesToIDs(categories)
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

func (s *service) mapIDsToActors(ids []string) []*Actor {
	actors := make([]*Actor, len(ids))
	for i, id := range ids {
		actors[i] = &Actor{ID: id}
	}
	return actors
}

func (s *service) mapIDsToCategories(ids []string) []*Category {
	categories := make([]*Category, len(ids))
	for i, id := range ids {
		categories[i] = &Category{ID: id}
	}
	return categories
}

func (s *service) mapActorsToIDs(actors []*Actor) []string {
	ids := make([]string, len(actors))
	for i, actor := range actors {
		ids[i] = actor.ID
	}
	return ids
}

func (s *service) mapCategoriesToIDs(categories []*Category) []string {
	ids := make([]string, len(categories))
	for i, category := range categories {
		ids[i] = category.ID
	}
	return ids
}
