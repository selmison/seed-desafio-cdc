package catalog

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// Book represents a single book.
type Book struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Title      string `gorm:"unique" validate:"required,not_blank"`
	Synopsis   string `validate:"required,not_blank,max=500"`
	Summary    string
	Price      float32     `validate:"required,gte=20"`
	Pages      int         `validate:"required,gte=100"`
	Isbn       string      `gorm:"unique" validate:"required,not_blank"`
	Issue      string      `validate:"required,not_blank"`
	Categories []*Category `gorm:"many2many:book_categories;"`
	Actors     []*Actor    `gorm:"many2many:actor_books;"`
}

func (a *Book) Validate() error {
	err := coreDomain.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		return catalogGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].Namespace(), coreDomain.ErrIsNotValid),
		)
	}
	return nil
}
