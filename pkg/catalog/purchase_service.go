package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreatePurchase implements create_customer.
func (s *service) CreatePurchase(ctx context.Context, dto *catalogGen.CreatePurchaseDTO) (res *catalogGen.PurchaseDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_purchase")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	customer := mapCreateCustomerDTOToCustomer(*dto.Customer)
	cart := mapCreateCartDTOToCart(*dto.Cart)
	purchase := Purchase{
		ID:         uuid.New().String(),
		CustomerID: customer.ID,
		Customer:   customer,
		CartID:     cart.ID,
		Cart:       cart,
	}
	if err := s.repo.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&customer).Error; err != nil {
			return err
		}
		if err := tx.Create(&purchase).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &catalogGen.PurchaseDTO{
		ID:       purchase.ID,
		Customer: mapCustomerToCustomerDTO(customer),
		Cart:     mapCartToCartDTO(cart),
	}, nil
}
