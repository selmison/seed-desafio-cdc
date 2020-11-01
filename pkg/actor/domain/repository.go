package domain

//go:generate mockgen -destination=./mock/repository.go -package=mock . Repository

import (
	"context"

	"github.com/selmison/seed-desafio-cdc/gen/actors"
)

// Repository persists Actors.
type Repository interface {
	// Store stores an Actor.
	Store(ctx context.Context, Actor Actor) (res *actors.ActorDTO, err error)
}
