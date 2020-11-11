package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Coupon represents a single actor.
type Coupon struct {
	gorm.Model
	ID       string  `gorm:"primaryKey"`
	Code     string  `gorm:"unique" validate:"required,not_blank"`
	Discount float32 `validate:"required,gt=0"`
	Validity string  `validate:"required,not_blank"`
	Carts    []*Cart
}

func (a *Coupon) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
