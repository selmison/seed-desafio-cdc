package catalog

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

// CreatePurchase implements create_customer.
func (s *service) CreatePurchase(ctx context.Context, dto *catalogGen.CreatePurchaseDTO) (res *catalogGen.PurchaseDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_purchase")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	customer := mapCreateCustomerDTOToCustomer(*dto.Customer)
	cart := mapCreateCartDTOToCart(*dto.Cart)
	total, totalWithCoupon, err := s.calculateTotals(cart)
	if err != nil {
		return nil, err
	}
	cart.total = total
	cart.totalWithCoupon = totalWithCoupon
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

// ShowPurchase implements show_purchase.
func (s *service) ShowPurchase(_ context.Context, dto *catalogGen.ShowByIDDTO) (res *catalogGen.PurchaseDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.show_purchase")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	purchase := Purchase{}
	result := s.repo.Where("id = ?", dto.ID).First(&purchase)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, catalogGen.MakeNotFound(coreDomain.ErrNotFound)
		}
		return nil, result.Error
	}
	customer := Customer{}
	if err := s.repo.Model(&purchase).Association("Customer").Find(&customer); err != nil {
		return nil, err
	}
	cart := Cart{}
	if err := s.repo.Model(&purchase).Association("Cart").Find(&cart); err != nil {
		return nil, err
	}
	items := []*Item{}
	if err := s.repo.Model(&cart).Association("Items").Find(&items); err != nil {
		return nil, err
	}
	cart.Items = items
	total, totalWithCoupon, err := s.calculateTotals(cart)
	if err != nil {
		return nil, err
	}
	cart.total = total
	cart.totalWithCoupon = totalWithCoupon
	return &catalogGen.PurchaseDTO{
		ID:       purchase.ID,
		Customer: mapCustomerToCustomerDTO(customer),
		Cart:     mapCartToCartDTO(cart),
	}, nil
}

func (s *service) calculateTotals(cart Cart) (total, totalWithCoupon float32, err error) {
	for _, item := range cart.Items {
		book := Book{}
		result := s.repo.Where("id = ?", item.BookID).First(&book)
		if result.Error != nil {
			return 0, 0, result.Error
		}
		total = total + float32(item.Amount)*book.Price
	}
	if cart.CouponID.Valid {
		coupon := Coupon{}
		result := s.repo.Where("id = ?", cart.CouponID).First(&coupon)
		if result.Error != nil {
			return 0, 0, nil
		}
		if result.RowsAffected > 0 {
			totalWithCoupon = total * (1 - coupon.Discount)
		}
	}
	return
}
