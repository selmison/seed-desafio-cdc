package infra

import (
	"context"
	"fmt"
	"testing"

	actorDomain "github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
)

func SetupActorTestCase(_ *testing.T, fakes []actorDomain.Actor) (func(t *testing.T), actorDomain.Repository, error) {
	repo := NewActorRepository()
	for _, actor := range fakes {
		if err := repo.Store(context.Background(), actor); err != nil {
			return nil, nil, fmt.Errorf("insert Actor: %s", err)
		}
	}
	return func(t *testing.T) {
		_ = repo.DeleteAll(nil)
	}, repo, nil
}
