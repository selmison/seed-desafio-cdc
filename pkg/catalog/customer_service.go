package catalog

import (
	"context"
	"fmt"
	"log"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

// CreateCustomer implements create_customer.
func (s *service) CreateCustomer(ctx context.Context, dto *catalogGen.CreateCustomerDTO) (res *catalogGen.CustomerDTO, err error) {
	if err := s.logger.Log("info", fmt.Sprintf("catalog.create_customer")); err != nil {
		log.Printf("kit/log error: %v\n", err)
	}
	customer := mapCreateCustomerDTOToCustomer(*dto)
	if err := customer.Validate(); err != nil {
		return nil, err
	}
	if result := s.repo.Create(&customer); result.Error != nil {
		return nil, result.Error
	}
	return &catalogGen.CustomerDTO{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Document:  customer.Document,
		Address: &catalogGen.AddressDTO{
			Address:    customer.Address.Address,
			Complement: customer.Address.Complement,
			City:       customer.Address.City,
			CountryID:  customer.Address.CountryID,
			StateID:    customer.Address.StateID,
			Cep:        customer.Address.Cep,
		},
		Phone:   customer.Phone,
		CartIds: mapCartsToIDs(customer.Carts),
	}, nil
}
