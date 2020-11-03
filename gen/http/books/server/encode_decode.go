// Code generated by goa v3.2.4, DO NOT EDIT.
//
// books HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"context"
	"io"
	"net/http"

	books "github.com/selmison/seed-desafio-cdc/gen/books"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateBookResponse returns an encoder for responses returned by the
// books create_book endpoint.
func EncodeCreateBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*books.BookDTO)
		enc := encoder(ctx, w)
		body := NewCreateBookResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateBookRequest returns a decoder for requests sent to the books
// create_book endpoint.
func DecodeCreateBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateBookRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateBookRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateBookDTO(&body)

		return payload, nil
	}
}

// EncodeCreateBookError returns an encoder for errors returned by the
// create_book books endpoint.
func EncodeCreateBookError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid_fields":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateBookInvalidFieldsResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_fields")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
