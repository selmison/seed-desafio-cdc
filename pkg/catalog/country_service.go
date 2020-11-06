package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreateCountry implements create_country.
func (s *service) CreateCountry(_ context.Context, dto *catalogGen.CreateCountryDTO) (res *catalogGen.CountryDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_country")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	country := Country{
		Model: gorm.Model{},
		ID:    uuid.New().String(),
		Name:  dto.Name,
	}
	if err := country.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&country); result.Error != nil {
		return nil, result.Error
	}
	return &catalogGen.CountryDTO{
		ID:   country.ID,
		Name: country.Name,
	}, nil
}
