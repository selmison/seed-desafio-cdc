package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type Customer struct {
	gorm.Model
	ID        string  `gorm:"primaryKey"`
	FirstName string  `validate:"required,not_blank"`
	LastName  string  `validate:"required,not_blank"`
	Email     string  `gorm:"unique" validate:"required,not_blank"`
	Document  string  `validate:"required,not_blank"`
	Address   Address `gorm:"embedded"`
	Phone     string  `validate:"required,not_blank"`
}

func (c *Customer) Validate() error {
	if err := coreDomain.Validate.Struct(c); err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	if err := c.Address.Validate(); err != nil {
		return err
	}
	return nil
}

type Address struct {
	Address    string `validate:"required,not_blank"`
	Complement string `validate:"required,not_blank"`
	City       string `validate:"required,not_blank"`
	StateID    string `validate:"required,not_blank"`
	State      State  `validate:"-"`
	Cep        string `validate:"required,not_blank"`
}

func (a *Address) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
