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

// CreateActor implements create.
func (s *service) CreateActor(_ context.Context, dto *catalogGen.CreateActorDTO) (res *catalogGen.ActorDTO, err error) {
	actor := Actor{
		ID:          uuid.New().String(),
		Name:        strings.TrimSpace(dto.Name),
		Email:       strings.TrimSpace(dto.Email),
		Description: strings.TrimSpace(dto.Description),
	}
	if err := actor.Validate(); err != nil {
		return nil, err
	}
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	result := s.repo.Create(&actor)
	return &catalogGen.ActorDTO{
		ID:          actor.ID,
		Name:        actor.Name,
		Email:       actor.Email,
		Description: actor.Description,
		CreatedAt:   actor.CreatedAt.String(),
	}, result.Error
}

// ShowActor implements show_actor.
func (s *service) ShowActor(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.ActorDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.show_actor")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	actor := Actor{}
	result := s.repo.Where("id = ?", dto.ID).First(&actor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	return &catalogGen.ActorDTO{
		ID:          actor.ID,
		Name:        actor.Name,
		Email:       actor.Email,
		Description: actor.Description,
		CreatedAt:   actor.CreatedAt.String(),
	}, nil
}
