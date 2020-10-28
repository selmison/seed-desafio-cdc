package domain

//go:generate mockgen -destination=./mock/repository.go -package=mock . Repository

import (
	"context"
)

// Repository persists Actors.
type Repository interface {
	// Store stores an Actor.
	Store(ctx context.Context, Actor Actor) error

	// GetAll returns all Actors.
	GetAll(ctx context.Context) ([]Actor, error)

	// GetMany returns Actors.
	GetMany(ctx context.Context, ids []string) ([]Actor, error)

	// GetOne returns a single Actor by its Id.
	GetOne(ctx context.Context, id string) (Actor, error)

	// DeleteAll deletes all actors.
	DeleteAll(ctx context.Context) error

	// DeleteOne deletes a single Actor by its Id.
	DeleteOne(ctx context.Context, id string) error

	// UpdateOne updates a single Actor by its Id.
	UpdateOne(ctx context.Context, id string, updateActor UpdateActorDTO) error
}
