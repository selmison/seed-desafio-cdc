package catalog

import (
	"context"
	"fmt"
	"log"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreateCart implements create_cart.
func (s *service) CreateCart(_ context.Context, dto *catalogGen.CreateCartDTO) (res *catalogGen.CartDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_cart")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	cart := mapCreateCartDTOToCart(*dto)
	if err := cart.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&cart); result.Error != nil {
		return nil, result.Error
	}
	return &catalogGen.CartDTO{
		ID:    cart.ID,
		Total: cart.Total(),
		Items: dto.Items,
	}, nil
}
