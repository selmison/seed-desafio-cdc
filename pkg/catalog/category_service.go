package catalog

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// CreateCategory implements create.
func (s *service) CreateCategory(_ context.Context, dto *catalogGen.CreateCategoryDTO) (res *catalogGen.CategoryDTO, err error) {
	category := Category{
		ID:   uuid.New().String(),
		Name: strings.TrimSpace(dto.Name),
	}
	if err := category.Validate(); err != nil {
		return nil, err
	}
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	result := s.repo.Create(&category)
	return &catalogGen.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}, result.Error
}

// ShowCategory implements show_category.
func (s *service) ShowCategory(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.CategoryDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.show_category")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	category := &Category{}
	result := s.repo.Where("id = ?", dto.ID).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	return &catalogGen.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}
