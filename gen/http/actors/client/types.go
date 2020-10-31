// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors HTTP client types
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"unicode/utf8"

	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "actors" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	Name        string `form:"name" json:"name" xml:"name"`
	EMail       string `form:"e-mail" json:"e-mail" xml:"e-mail"`
	Description string `form:"description" json:"description" xml:"description"`
}

// CreateResponseBody is the type of the "actors" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	ID          *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	EMail       *string `form:"e-mail,omitempty" json:"e-mail,omitempty" xml:"e-mail,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	CreatedAt   *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "actors" service.
func NewCreateRequestBody(p *actors.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:        p.Name,
		EMail:       p.EMail,
		Description: p.Description,
	}
	return body
}

// NewCreateActorPayloadCreated builds a "actors" service "create" endpoint
// result from a HTTP "Created" response.
func NewCreateActorPayloadCreated(body *CreateResponseBody) *actors.ActorPayload {
	v := &actors.ActorPayload{
		ID:          *body.ID,
		Name:        *body.Name,
		EMail:       *body.EMail,
		Description: *body.Description,
		CreatedAt:   *body.CreatedAt,
	}

	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.EMail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("e-mail", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 400, false))
		}
	}
	return
}
