// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"context"
	"io"
	"net/http"

	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the actors
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*actors.ActorPayload)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the actors create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(&body)

		return payload, nil
	}
}
