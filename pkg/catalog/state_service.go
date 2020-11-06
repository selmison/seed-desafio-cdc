package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreateState implements create_state.
func (s *service) CreateState(ctx context.Context, dto *catalogGen.CreateStateDTO) (res *catalogGen.StateDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_state")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	state := State{
		ID:        uuid.New().String(),
		Name:      dto.Name,
		CountryID: dto.CountryID,
	}
	if err := state.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&state); result.Error != nil {
		return nil, result.Error
	}
	return &catalogGen.StateDTO{
		ID:        state.ID,
		Name:      state.Name,
		CountryID: state.CountryID,
	}, nil

}
