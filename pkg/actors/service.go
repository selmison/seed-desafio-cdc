package actors

import (
	"context"
	"fmt"
	"log"
	"strings"

	kitLog "github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/actors"
)

type service struct {
	repo   *gorm.DB
	logger kitLog.Logger
}

// NewService returns a new Actors Service.
func NewService(repo *gorm.DB, logger kitLog.Logger) actors.Service {
	return &service{repo, logger}
}

// CreateActor implements create.
func (s *service) CreateActor(_ context.Context, dto *actors.CreateActorDTO) (res *actors.ActorDTO, err error) {
	actor := Actor{
		ID:          uuid.New().String(),
		Name:        strings.TrimSpace(dto.Name),
		Email:       strings.TrimSpace(dto.EMail),
		Description: strings.TrimSpace(dto.Description),
	}
	if err := actor.Validate(); err != nil {
		return nil, err
	}
	if err := s.logger.Log("info", fmt.Sprintf("actors.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	result := s.repo.Create(&actor)
	return &actors.ActorDTO{
		ID:          actor.ID,
		Name:        actor.Name,
		EMail:       actor.Email,
		Description: actor.Description,
		CreatedAt:   actor.CreatedAt.String(),
	}, result.Error
}
