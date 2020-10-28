package service

import (
	"context"
	"strings"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

type validationMiddleware struct {
	next Service
}

// NewValidationMiddleware returns a service Middleware.
func NewValidationMiddleware() Middleware {
	return func(next Service) Service {
		return validationMiddleware{next: next}
	}
}

func (m validationMiddleware) Create(ctx context.Context, actor domain.Actor) error {
	return m.next.Create(ctx, actor)
}

func (m validationMiddleware) Destroy(ctx context.Context, id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return coreDomain.ErrIdCouldNotBeEmpty
	}
	return m.next.Destroy(ctx, id)
}

func (m validationMiddleware) List(ctx context.Context) ([]domain.Actor, error) {
	return m.next.List(ctx)
}

func (m validationMiddleware) Show(ctx context.Context, id string) (domain.Actor, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return domain.Actor{}, coreDomain.ErrIdCouldNotBeEmpty
	}
	return m.next.Show(ctx, id)
}

func (m validationMiddleware) Update(ctx context.Context, id string, updateActor domain.UpdateActorDTO) error {
	if updateActor.IsEmpty() {
		return nil
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return coreDomain.ErrIdCouldNotBeEmpty
	}
	return m.next.Update(ctx, id, updateActor)
}
