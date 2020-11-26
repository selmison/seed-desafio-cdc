// Code generated by goa v3.2.4, DO NOT EDIT.
//
// catalog go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/selmison/seed-desafio-cdc/gen/http/catalog/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateActorResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_actor responses.
func EncodeCreateActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateActorResponse(encoder)
}

// DecodeCreateActorRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_actor requests.
func DecodeCreateActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateActorRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateActorError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_actor endpoint.
func EncodeCreateActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateActorError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListActorsResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog list_actors responses.
func EncodeListActorsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListActorsResponse(encoder)
}

// EncodeShowActorResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_actor responses.
func EncodeShowActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowActorResponse(encoder)
}

// DecodeShowActorRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_actor requests.
func DecodeShowActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowActorRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowActorError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_actor endpoint.
func EncodeShowActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowActorError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateBookResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_book responses.
func EncodeCreateBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateBookResponse(encoder)
}

// DecodeCreateBookRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_book requests.
func DecodeCreateBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateBookRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateBookError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_book endpoint.
func EncodeCreateBookError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateBookError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListBooksResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog list_books responses.
func EncodeListBooksResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListBooksResponse(encoder)
}

// EncodeShowBookResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_book responses.
func EncodeShowBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowBookResponse(encoder)
}

// DecodeShowBookRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_book requests.
func DecodeShowBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowBookRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowBookError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_book endpoint.
func EncodeShowBookError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowBookError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateCartResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_cart responses.
func EncodeCreateCartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateCartResponse(encoder)
}

// DecodeCreateCartRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_cart requests.
func DecodeCreateCartRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateCartRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateCartError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_cart endpoint.
func EncodeCreateCartError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateCartError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateCategoryResponse returns a go-kit EncodeResponseFunc suitable
// for encoding catalog create_category responses.
func EncodeCreateCategoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateCategoryResponse(encoder)
}

// DecodeCreateCategoryRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_category requests.
func DecodeCreateCategoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateCategoryRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateCategoryError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_category endpoint.
func EncodeCreateCategoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateCategoryError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListCategoriesResponse returns a go-kit EncodeResponseFunc suitable
// for encoding catalog list_categories responses.
func EncodeListCategoriesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListCategoriesResponse(encoder)
}

// EncodeShowCategoryResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_category responses.
func EncodeShowCategoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowCategoryResponse(encoder)
}

// DecodeShowCategoryRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_category requests.
func DecodeShowCategoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowCategoryRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowCategoryError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_category endpoint.
func EncodeShowCategoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowCategoryError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateCountryResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_country responses.
func EncodeCreateCountryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateCountryResponse(encoder)
}

// DecodeCreateCountryRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_country requests.
func DecodeCreateCountryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateCountryRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateCountryError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_country endpoint.
func EncodeCreateCountryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateCountryError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListCountriesResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog list_countries responses.
func EncodeListCountriesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListCountriesResponse(encoder)
}

// EncodeShowCountryResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_country responses.
func EncodeShowCountryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowCountryResponse(encoder)
}

// DecodeShowCountryRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_country requests.
func DecodeShowCountryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowCountryRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowCountryError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_country endpoint.
func EncodeShowCountryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowCountryError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeApplyCouponResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog apply_coupon responses.
func EncodeApplyCouponResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeApplyCouponResponse(encoder)
}

// DecodeApplyCouponRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog apply_coupon requests.
func DecodeApplyCouponRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeApplyCouponRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeApplyCouponError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog apply_coupon endpoint.
func EncodeApplyCouponError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeApplyCouponError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateCouponResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_coupon responses.
func EncodeCreateCouponResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateCouponResponse(encoder)
}

// DecodeCreateCouponRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_coupon requests.
func DecodeCreateCouponRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateCouponRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateCouponError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_coupon endpoint.
func EncodeCreateCouponError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateCouponError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateCustomerResponse returns a go-kit EncodeResponseFunc suitable
// for encoding catalog create_customer responses.
func EncodeCreateCustomerResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateCustomerResponse(encoder)
}

// DecodeCreateCustomerRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_customer requests.
func DecodeCreateCustomerRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateCustomerRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateCustomerError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_customer endpoint.
func EncodeCreateCustomerError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateCustomerError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreatePurchaseResponse returns a go-kit EncodeResponseFunc suitable
// for encoding catalog create_purchase responses.
func EncodeCreatePurchaseResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreatePurchaseResponse(encoder)
}

// DecodeCreatePurchaseRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_purchase requests.
func DecodeCreatePurchaseRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreatePurchaseRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreatePurchaseError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_purchase endpoint.
func EncodeCreatePurchaseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreatePurchaseError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeShowPurchaseResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_purchase responses.
func EncodeShowPurchaseResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowPurchaseResponse(encoder)
}

// DecodeShowPurchaseRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_purchase requests.
func DecodeShowPurchaseRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowPurchaseRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowPurchaseError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_purchase endpoint.
func EncodeShowPurchaseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowPurchaseError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreateStateResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog create_state responses.
func EncodeCreateStateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateStateResponse(encoder)
}

// DecodeCreateStateRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog create_state requests.
func DecodeCreateStateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateStateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateStateError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog create_state endpoint.
func EncodeCreateStateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateStateError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListStatesResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog list_states responses.
func EncodeListStatesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListStatesResponse(encoder)
}

// EncodeShowStateResponse returns a go-kit EncodeResponseFunc suitable for
// encoding catalog show_state responses.
func EncodeShowStateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowStateResponse(encoder)
}

// DecodeShowStateRequest returns a go-kit DecodeRequestFunc suitable for
// decoding catalog show_state requests.
func DecodeShowStateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowStateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowStateError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the catalog show_state endpoint.
func EncodeShowStateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowStateError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
