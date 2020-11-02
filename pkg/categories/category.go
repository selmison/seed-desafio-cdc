package categories

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type Category struct {
	gorm.Model
	ID   string `gorm:"primarykey"`
	Name string `gorm:"unique" validate:"required,not_blank"`
}

func (c *Category) Validate() error {
	err := coreDomain.Validate.Struct(c)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValidated)
	}
	return nil
}
