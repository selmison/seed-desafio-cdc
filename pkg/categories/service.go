package categories

import (
	"context"
	"fmt"
	"log"
	"strings"

	kitLog "github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/categories"
)

type service struct {
	repo   *gorm.DB
	logger kitLog.Logger
}

// NewService returns a new Categories Service.
func NewService(repo *gorm.DB, logger kitLog.Logger) categories.Service {
	return &service{repo, logger}
}

// CreateCategory implements create.
func (s *service) CreateCategory(_ context.Context, dto *categories.CreateCategoryDTO) (res *categories.CategoryDTO, err error) {
	actor := Category{
		ID:   uuid.New().String(),
		Name: strings.TrimSpace(dto.Name),
	}
	if err := actor.Validate(); err != nil {
		return nil, err
	}
	if err := s.logger.Log("info", fmt.Sprintf("categories.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	result := s.repo.Create(&actor)
	return &categories.CategoryDTO{
		ID:   actor.ID,
		Name: actor.Name,
	}, result.Error
}
