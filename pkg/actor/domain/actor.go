package domain

import (
	"time"

	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Actor represents a single actor.
type Actor struct {
	Id          domain.Id
	Name        Name
	Email       Email
	Description Desc
	createdAt   time.Time
}

func (a *Actor) CreatAt() time.Time {
	return a.createdAt
}
