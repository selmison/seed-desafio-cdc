package infra

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

var (
	actorsOnce sync.Once
	repo       *actorRepository
)

// actorRepository keeps actors in the memory.
// Use it in tests or for development/demo purposes.
type actorRepository struct {
	actors     map[string]domain.Actor
	actorsOnce sync.Once
	mu         sync.RWMutex
}

// NewActorRepository returns a new in-memory actor actorRepository.
func NewActorRepository() domain.Repository {
	actorsOnce.Do(func() {
		repo = &actorRepository{}
		repo.init()
	})
	return repo
}

func (r *actorRepository) init() {
	r.actorsOnce.Do(func() {
		r.actors = make(map[string]domain.Actor)
	})
}

// Store stores an actor.
func (r *actorRepository) Store(_ context.Context, actor domain.Actor) error {
	r.init()
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, value := range r.actors {
		if value.Id == actor.Id {
			return fmt.Errorf("id '%s' %w", actor.Id.String(), coreDomain.ErrAlreadyExists)
		}
	}
	r.actors[actor.Id.String()] = actor
	return nil
}

// GetAll returns all actors.
func (r *actorRepository) GetAll(ctx context.Context) ([]domain.Actor, error) {
	return r.GetMany(ctx, nil)
}

// GetMany returns a list of actors filtered by ids.
func (r *actorRepository) GetMany(_ context.Context, ids []string) ([]domain.Actor, error) {
	r.init()
	r.mu.RLock()
	defer r.mu.RUnlock()
	var actors []domain.Actor
	length := len(ids)
	if length == 0 {
		length = len(r.actors)
		for _, k := range r.actors {
			ids = append(ids, k.Id.String())
		}
	}
	// This makes sure actors are always returned in the same, sorted order
	sort.Strings(ids)
	for _, id := range ids {
		if !compare(r.actors[id], domain.Actor{}) {
			actors = append(actors, r.actors[id])
		}
	}
	return actors, nil
}

// GetOne returns a single actor by its Id.
func (r *actorRepository) GetOne(_ context.Context, id string) (domain.Actor, error) {
	r.init()
	r.mu.RLock()
	defer r.mu.RUnlock()
	actor, ok := r.actors[id]
	if !ok {
		return actor, fmt.Errorf("%s: %w", id, coreDomain.ErrNotFound)
	}
	return actor, nil
}

// DeleteAll deletes all actors from the actorRepository.
func (r *actorRepository) DeleteAll(_ context.Context) error {
	r.init()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.actors = make(map[string]domain.Actor)
	return nil
}

// DeleteOne deletes a single actor by its Id.
func (r *actorRepository) DeleteOne(_ context.Context, id string) error {
	r.init()
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.actors[id]
	if !ok {
		return fmt.Errorf("%s: %w", id, coreDomain.ErrNotFound)
	}
	delete(r.actors, id)
	return nil
}

// UpdateOne updates a single actor by its Id.
func (r *actorRepository) UpdateOne(_ context.Context, id string, updateActor domain.UpdateActorDTO) error {
	r.init()
	r.mu.Lock()
	defer r.mu.Unlock()
	actor, ok := r.actors[id]
	if !ok {
		return fmt.Errorf("%s: %w", id, coreDomain.ErrNotFound)
	}
	actor, err := update(actor, updateActor)
	if err != nil {
		return err
	}
	r.actors[id] = actor
	return nil
}

func update(actor domain.Actor, updateActor domain.UpdateActorDTO) (domain.Actor, error) {
	name := actor.Name
	if updateActor.Name != nil {
		name = *updateActor.Name
	}
	email := actor.Email
	if updateActor.Email != nil {
		email = *updateActor.Email
	}
	desc := actor.Description
	if updateActor.Description != nil {
		desc = *updateActor.Description
	}
	return domain.Actor{
		Id:          actor.Id,
		Name:        name,
		Email:       email,
		Description: desc,
	}, nil
}

func compare(a, b domain.Actor) bool {
	if &a == &b {
		return true
	}
	if a.Id != b.Id {
		return false
	}
	if a.Name != b.Name {
		return false
	}
	if b.Email != b.Email {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	return true
}
