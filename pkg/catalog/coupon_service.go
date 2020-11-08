package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

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
