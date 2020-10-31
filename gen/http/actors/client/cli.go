// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors HTTP client CLI support package
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the actors create endpoint from
// CLI flags.
func BuildCreatePayload(actorsCreateBody string) (*actors.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(actorsCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"w4k\",\n      \"e-mail\": \"Aliquam asperiores iusto.\",\n      \"name\": \"Enim itaque quod cupiditate.\"\n   }'")
		}
		if utf8.RuneCountInString(body.Description) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 400, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &actors.CreatePayload{
		Name:        body.Name,
		EMail:       body.EMail,
		Description: body.Description,
	}

	return v, nil
}
