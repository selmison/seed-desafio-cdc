package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// ApplyCoupon implements apply.
func (s *service) ApplyCoupon(_ context.Context, dto *catalogGen.ApplyCouponDTO) error {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.apply")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	coupon := Coupon{}
	if result := s.repo.Where("code = ?", dto.Code).First(&coupon); result.Error != nil {
		return result.Error
	}
	purchase := Purchase{}
	if result := s.repo.Where("id = ?", dto.PurchaseID).First(&purchase); result.Error != nil {
		return result.Error
	}
	cart := Cart{}
	if err := s.repo.Model(&purchase).Association("Cart").Find(&cart); err != nil {
		return err
	}
	if result := s.repo.Model(&cart).Update("coupon_id", coupon.ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateCoupon implements create.
func (s *service) CreateCoupon(_ context.Context, dto *catalogGen.CreateCouponDTO) (res *catalogGen.CouponDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	coupon := Coupon{
		ID:       uuid.New().String(),
		Code:     dto.Code,
		Discount: dto.Discount,
		Validity: dto.Validity,
	}
	if err := coupon.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&coupon); result.Error != nil {
		return nil, result.Error
	}
	return &catalogGen.CouponDTO{
		ID:       coupon.ID,
		Code:     coupon.Code,
		Discount: coupon.Discount,
		Validity: coupon.Validity,
	}, nil
}
