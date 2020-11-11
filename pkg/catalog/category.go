package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type Category struct {
	gorm.Model
	ID    string  `gorm:"primaryKey"`
	Name  string  `gorm:"unique" validate:"required,not_blank"`
	Books []*Book `gorm:"many2many:book_categories;"`
}

func (c *Category) Validate() error {
	err := coreDomain.Validate.Struct(c)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
