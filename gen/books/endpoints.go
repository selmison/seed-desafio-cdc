// Code generated by goa v3.2.4, DO NOT EDIT.
//
// books endpoints
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package books

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "books" service endpoints.
type Endpoints struct {
	CreateBook endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "books" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateBook: NewCreateBookEndpoint(s),
	}
}

// Use applies the given middleware to all the "books" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.CreateBook = m(e.CreateBook)
}

// NewCreateBookEndpoint returns an endpoint function that calls the method
// "create_book" of service "books".
func NewCreateBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateBookDTO)
		return s.CreateBook(ctx, p)
	}
}
