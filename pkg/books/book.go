package books

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/books"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Book represents a single book.
type Book struct {
	gorm.Model
	ID         string `gorm:"primarykey"`
	Title      string `gorm:"unique" validate:"required,not_blank"`
	Synopsis   string `validate:"required,not_blank,max=500"`
	Summary    string
	Price      float32 `validate:"required,gte=20"`
	Pages      int     `validate:"required,gte=100"`
	Isbn       string  `gorm:"unique" validate:"required,not_blank"`
	Issue      string  `validate:"required,not_blank"`
	CategoryID string  `validate:"required,not_blank"`
	ActorID    string  `validate:"required,not_blank"`
}

func (a *Book) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		err := books.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), coreDomain.ErrIsNotValid),
		)
		return err
	}
	return nil
}
