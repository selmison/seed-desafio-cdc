// Code generated by goa v3.2.4, DO NOT EDIT.
//
// categories go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/selmison/seed-desafio-cdc/gen/http/categories/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateCategoryRequest returns a go-kit EncodeRequestFunc suitable for
// encoding categories create_category requests.
func EncodeCreateCategoryRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateCategoryRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateCategoryResponse returns a go-kit DecodeResponseFunc suitable
// for decoding categories create_category responses.
func DecodeCreateCategoryResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateCategoryResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
