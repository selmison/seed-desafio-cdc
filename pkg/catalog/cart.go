package catalog

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Cart represents a single cart.
type Cart struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Items    []*Item
	CouponID sql.NullString
}

func (c *Cart) Validate() error {
	if err := coreDomain.Validate.Struct(c); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}

func (c *Cart) Total() float32 {
	return 0
}

type Item struct {
	gorm.Model
	ID     string `gorm:"primaryKey"`
	BookID string `validate:"required,not_blank"`
	Amount int32  `validate:"required,gt=0"`
	CartID string
}

func (i *Item) Validate() error {
	if err := coreDomain.Validate.Struct(i); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
