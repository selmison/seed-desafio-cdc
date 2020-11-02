package actors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/actors"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Actor represents a single actor.
type Actor struct {
	gorm.Model
	ID          string `gorm:"primarykey"`
	Name        string `validate:"required,not_blank"`
	Email       string `gorm:"unique" validate:"required,email"`
	Description string `validate:"required,not_blank,max=400"`
}

func (a *Actor) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		err := actors.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValid),
		)
		return err
	}
	return nil
}
