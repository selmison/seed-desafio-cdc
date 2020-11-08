package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type Cart struct {
	gorm.Model
	ID         string  `gorm:"primarykey"`
	Total      float32 `validate:"required"`
	Items      []*Item
	CustomerID string `validate:"required,not_blank"`
}

func (c *Cart) Validate() error {
	if err := coreDomain.Validate.Struct(c); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValid),
		)
	}
	for _, item := range c.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Item struct {
	BookID string `validate:"required,not_blank"`
	Amount int32
	CartID string
}

func (i *Item) Validate() error {
	if err := coreDomain.Validate.Struct(i); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
