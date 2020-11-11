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

// ListStates implements list_actor.
func (s *service) ListStates(context.Context) (res []*catalogGen.StateDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("actors.list_actors")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	var actors []*State
	result := s.repo.Find(&actors)
	if result.Error != nil {
		return nil, result.Error
	}
	actorDTOS := make([]*catalogGen.StateDTO, len(actors))
	for i, state := range actors {
		actorDTOS[i] = &catalogGen.StateDTO{
			ID:        state.ID,
			Name:      state.Name,
			CountryID: state.CountryID,
		}
	}
	return actorDTOS, nil
}

// ShowState implements show_actor.
func (s *service) ShowState(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.StateDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.show_actor")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	state := State{}
	result := s.repo.Where("id = ?", dto.ID).First(&state)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	return &catalogGen.StateDTO{
		ID:        state.ID,
		Name:      state.Name,
		CountryID: state.CountryID,
	}, nil
}
