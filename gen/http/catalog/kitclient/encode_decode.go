// Code generated by goa v3.2.4, DO NOT EDIT.
//
// catalog go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/selmison/seed-desafio-cdc/gen/http/catalog/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateActorRequest returns a go-kit EncodeRequestFunc suitable for
// encoding catalog create_actor requests.
func EncodeCreateActorRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateActorRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateActorResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog create_actor responses.
func DecodeCreateActorResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateActorResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// DecodeShowActorResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog show_actor responses.
func DecodeShowActorResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeShowActorResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeCreateBookRequest returns a go-kit EncodeRequestFunc suitable for
// encoding catalog create_book requests.
func EncodeCreateBookRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateBookRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateBookResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog create_book responses.
func DecodeCreateBookResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateBookResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// DecodeListBooksResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog list_books responses.
func DecodeListBooksResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeListBooksResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// DecodeShowBookResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog show_book responses.
func DecodeShowBookResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeShowBookResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeCreateCategoryRequest returns a go-kit EncodeRequestFunc suitable for
// encoding catalog create_category requests.
func EncodeCreateCategoryRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateCategoryRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateCategoryResponse returns a go-kit DecodeResponseFunc suitable
// for decoding catalog create_category responses.
func DecodeCreateCategoryResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateCategoryResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// DecodeShowCategoryResponse returns a go-kit DecodeResponseFunc suitable for
// decoding catalog show_category responses.
func DecodeShowCategoryResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeShowCategoryResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
