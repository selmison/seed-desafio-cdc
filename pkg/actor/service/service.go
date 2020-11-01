package service

import (
	"context"
	"fmt"

	kitLog "github.com/go-kit/kit/log"

	"github.com/selmison/seed-desafio-cdc/gen/actors"
	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type service struct {
	repo   domain.Repository
	logger kitLog.Logger
}

// NewService returns a new Actors Service.
func NewService(repo domain.Repository, logger kitLog.Logger) actors.Service {
	return &service{repo, logger}
}

// Create implements create.
func (s *service) Create(ctx context.Context, dto *actors.CreateActorDTO) (res *actors.ActorDTO, err error) {
	id, err := coreDomain.GenerateID()
	if err != nil {
		return nil, err
	}
	name, err := domain.NewName(dto.Name)
	if err != nil {
		return nil, err
	}
	email, err := domain.NewEmail(dto.EMail)
	if err != nil {
		return nil, err
	}
	desc, err := domain.NewDesc(dto.Description)
	if err != nil {
		return nil, err
	}
	createdAt := domain.GenerateTime()
	actor := domain.Actor{
		ID:          id,
		Name:        name,
		Email:       email,
		Description: desc,
		CreatedAt:   createdAt,
	}
	s.logger.Log("info", fmt.Sprintf("actors.create"))
	return s.repo.Store(ctx, actor)
}
