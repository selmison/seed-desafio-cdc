package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Actor represents a single actor.
type Actor struct {
	gorm.Model
	ID          string `gorm:"primarykey"`
	Name        string `validate:"required,not_blank"`
	Email       string `validate:"required,email"`
	Description string `validate:"required,not_blank,max=400"`
}

func (v *Actor) Validate() error {
	err := coreDomain.Validate.Struct(v)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValidated)
	}
	return nil
}
