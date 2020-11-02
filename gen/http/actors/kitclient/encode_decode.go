// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/selmison/seed-desafio-cdc/gen/http/actors/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateRequest returns a go-kit EncodeRequestFunc suitable for encoding
// actors create requests.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateResponse returns a go-kit DecodeResponseFunc suitable for
// decoding actors create responses.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}