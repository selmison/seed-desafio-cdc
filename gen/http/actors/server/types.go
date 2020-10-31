// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors HTTP server types
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"unicode/utf8"

	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "actors" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	EMail       *string `form:"e-mail,omitempty" json:"e-mail,omitempty" xml:"e-mail,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// CreateResponseBody is the type of the "actors" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	ID          string `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	EMail       string `form:"e-mail" json:"e-mail" xml:"e-mail"`
	Description string `form:"description" json:"description" xml:"description"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "actors" service.
func NewCreateResponseBody(res *actors.ActorPayload) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		EMail:       res.EMail,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}
	return body
}

// NewCreatePayload builds a actors service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody) *actors.CreatePayload {
	v := &actors.CreatePayload{
		Name:        *body.Name,
		EMail:       *body.EMail,
		Description: *body.Description,
	}

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.EMail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("e-mail", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 400, false))
		}
	}
	return
}
