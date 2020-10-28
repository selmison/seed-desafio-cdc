package domain

import (
	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Actor represents a single actor.
type Actor struct {
	Id          domain.Id
	Name        Name
	Email       Email
	Description Desc
	CreatedAt   Time
}
