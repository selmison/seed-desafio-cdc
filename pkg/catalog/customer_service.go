package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreateCustomer implements create_customer.
func (s *service) CreateCustomer(ctx context.Context, p *catalog.CreateCustomerDTO) (res *catalog.CustomerDTO, err error) {
	res = &catalog.CustomerDTO{}
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_customer")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	return
}
