package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Purchase represents a single purchase.
type Purchase struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	CustomerID string `gorm:"not null"`
	Customer   Customer
	CartID     string `gorm:"not null"`
	Cart       Cart
}

func (p *Purchase) Validate() error {
	if err := coreDomain.Validate.Struct(p); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
