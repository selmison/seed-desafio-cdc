package service

import (
	"context"

	kitLog "github.com/go-kit/kit/log"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
)

type Service interface {
	// Create creates a new actor.
	Create(ctx context.Context, actor domain.Actor) error

	// Destroy destroys a actor.
	Destroy(ctx context.Context, name string) error

	// List returns a list of actors.
	List(ctx context.Context) ([]domain.Actor, error)

	// Show returns the details of a actor.
	Show(ctx context.Context, id string) (domain.Actor, error)

	// Update updates an existing actor.
	Update(ctx context.Context, id string, updateActor domain.UpdateActorDTO) error
}

type service struct {
	repo domain.Repository
}

// NewService returns a new Service with all of the expected middlewares wired in.
func NewService(r domain.Repository, logger kitLog.Logger) Service {
	var svc Service
	{
		svc = service{repo: r}
		svc = NewValidationMiddleware()(svc)
		svc = NewLoggingMiddleware(logger)(svc)
	}
	return svc
}

func (svc service) Create(ctx context.Context, actor domain.Actor) error {
	return svc.repo.Store(ctx, actor)
}

func (svc service) Destroy(ctx context.Context, id string) error {
	return svc.repo.DeleteOne(ctx, id)
}

func (svc service) List(ctx context.Context) ([]domain.Actor, error) {
	actors, err := svc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (svc service) Show(ctx context.Context, id string) (domain.Actor, error) {
	actor, err := svc.repo.GetOne(ctx, id)
	if err != nil {
		return domain.Actor{}, err
	}
	return actor, nil
}

func (svc service) Update(ctx context.Context, id string, updateActor domain.UpdateActorDTO) error {
	err := svc.repo.UpdateOne(ctx, id, updateActor)
	if err != nil {
		return err
	}
	return nil
}
