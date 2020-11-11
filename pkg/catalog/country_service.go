package catalog

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
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

// ListCountries implements list_country.
func (s *service) ListCountries(context.Context) (res []*catalogGen.CountryDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("countries.list_countries")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	var countries []*Country
	result := s.repo.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	countryDTOS := make([]*catalogGen.CountryDTO, len(countries))
	for i, country := range countries {
		countryDTOS[i] = &catalogGen.CountryDTO{
			ID:   country.ID,
			Name: country.Name,
		}
	}
	return countryDTOS, nil
}

// ShowCountry implements show_country.
func (s *service) ShowCountry(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.CountryDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.show_country")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	country := &Country{}
	result := s.repo.Where("id = ?", dto.ID).First(&country)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	return &catalogGen.CountryDTO{
		ID:   country.ID,
		Name: country.Name,
	}, nil
}
