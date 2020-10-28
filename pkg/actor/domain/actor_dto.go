package domain

import (
	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// ActorDTO contains the details of a Actor.
type ActorDTO struct {
	Id          domain.Id `json:"id"`
	Name        Name      `json:"name"`
	Email       Email     `json:"e-mail"`
	Description Desc      `json:"description"`
	CreatedAt   Time      `json:"created_at"`
}