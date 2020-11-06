package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// State represents a single state.
type State struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	Name      string `validate:"required,not_blank"`
	CountryID string
}

func (a *State) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
