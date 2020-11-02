// Code generated by goa v3.2.4, DO NOT EDIT.
//
// categories HTTP client types
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	categories "github.com/selmison/seed-desafio-cdc/gen/categories"
	goa "goa.design/goa/v3/pkg"
)

// CreateCategoryRequestBody is the type of the "categories" service
// "create_category" endpoint HTTP request body.
type CreateCategoryRequestBody struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateCategoryResponseBody is the type of the "categories" service
// "create_category" endpoint HTTP response body.
type CreateCategoryResponseBody struct {
	ID   *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewCreateCategoryRequestBody builds the HTTP request body from the payload
// of the "create_category" endpoint of the "categories" service.
func NewCreateCategoryRequestBody(p *categories.CreateCategoryDTO) *CreateCategoryRequestBody {
	body := &CreateCategoryRequestBody{
		Name: p.Name,
	}
	return body
}

// NewCreateCategoryCategoryDTOCreated builds a "categories" service
// "create_category" endpoint result from a HTTP "Created" response.
func NewCreateCategoryCategoryDTOCreated(body *CreateCategoryResponseBody) *categories.CategoryDTO {
	v := &categories.CategoryDTO{
		ID:   *body.ID,
		Name: *body.Name,
	}

	return v
}

// ValidateCreateCategoryResponseBody runs the validations defined on
// create_category_response_body
func ValidateCreateCategoryResponseBody(body *CreateCategoryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
