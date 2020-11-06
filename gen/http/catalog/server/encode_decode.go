// Code generated by goa v3.2.4, DO NOT EDIT.
//
// catalog HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"context"
	"io"
	"net/http"

	catalog "github.com/selmison/seed-desafio-cdc/gen/catalog"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateActorResponse returns an encoder for responses returned by the
// catalog create_actor endpoint.
func EncodeCreateActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.ActorDTO)
		enc := encoder(ctx, w)
		body := NewCreateActorResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateActorRequest returns a decoder for requests sent to the catalog
// create_actor endpoint.
func DecodeCreateActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateActorRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateActorRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateActorDTO(&body)

		return payload, nil
	}
}

// EncodeCreateActorError returns an encoder for errors returned by the
// create_actor catalog endpoint.
func EncodeCreateActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewCreateActorInvalidFieldsResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_fields")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowActorResponse returns an encoder for responses returned by the
// catalog show_actor endpoint.
func EncodeShowActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.ActorDTO)
		enc := encoder(ctx, w)
		body := NewShowActorResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowActorRequest returns a decoder for requests sent to the catalog
// show_actor endpoint.
func DecodeShowActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewShowActorShowByIDDTO(id)

		return payload, nil
	}
}

// EncodeShowActorError returns an encoder for errors returned by the
// show_actor catalog endpoint.
func EncodeShowActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowActorNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateBookResponse returns an encoder for responses returned by the
// catalog create_book endpoint.
func EncodeCreateBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.BookDTO)
		enc := encoder(ctx, w)
		body := NewCreateBookResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateBookRequest returns a decoder for requests sent to the catalog
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
// create_book catalog endpoint.
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

// EncodeListBooksResponse returns an encoder for responses returned by the
// catalog list_books endpoint.
func EncodeListBooksResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*catalog.BookDTO)
		enc := encoder(ctx, w)
		body := NewListBooksResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeShowBookResponse returns an encoder for responses returned by the
// catalog show_book endpoint.
func EncodeShowBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.BookDTO)
		enc := encoder(ctx, w)
		body := NewShowBookResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowBookRequest returns a decoder for requests sent to the catalog
// show_book endpoint.
func DecodeShowBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewShowBookShowByIDDTO(id)

		return payload, nil
	}
}

// EncodeShowBookError returns an encoder for errors returned by the show_book
// catalog endpoint.
func EncodeShowBookError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowBookNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateCategoryResponse returns an encoder for responses returned by
// the catalog create_category endpoint.
func EncodeCreateCategoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.CategoryDTO)
		enc := encoder(ctx, w)
		body := NewCreateCategoryResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateCategoryRequest returns a decoder for requests sent to the
// catalog create_category endpoint.
func DecodeCreateCategoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateCategoryRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateCategoryRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateCategoryDTO(&body)

		return payload, nil
	}
}

// EncodeCreateCategoryError returns an encoder for errors returned by the
// create_category catalog endpoint.
func EncodeCreateCategoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewCreateCategoryInvalidFieldsResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_fields")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowCategoryResponse returns an encoder for responses returned by the
// catalog show_category endpoint.
func EncodeShowCategoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.CategoryDTO)
		enc := encoder(ctx, w)
		body := NewShowCategoryResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowCategoryRequest returns a decoder for requests sent to the catalog
// show_category endpoint.
func DecodeShowCategoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewShowCategoryShowByIDDTO(id)

		return payload, nil
	}
}

// EncodeShowCategoryError returns an encoder for errors returned by the
// show_category catalog endpoint.
func EncodeShowCategoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowCategoryNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateCountryResponse returns an encoder for responses returned by the
// catalog create_country endpoint.
func EncodeCreateCountryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.CountryDTO)
		enc := encoder(ctx, w)
		body := NewCreateCountryResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateCountryRequest returns a decoder for requests sent to the
// catalog create_country endpoint.
func DecodeCreateCountryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateCountryRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateCountryRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateCountryDTO(&body)

		return payload, nil
	}
}

// EncodeCreateCountryError returns an encoder for errors returned by the
// create_country catalog endpoint.
func EncodeCreateCountryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewCreateCountryInvalidFieldsResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_fields")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateStateResponse returns an encoder for responses returned by the
// catalog create_state endpoint.
func EncodeCreateStateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*catalog.StateDTO)
		enc := encoder(ctx, w)
		body := NewCreateStateResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateStateRequest returns a decoder for requests sent to the catalog
// create_state endpoint.
func DecodeCreateStateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateStateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateStateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateStateDTO(&body)

		return payload, nil
	}
}

// EncodeCreateStateError returns an encoder for errors returned by the
// create_state catalog endpoint.
func EncodeCreateStateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewCreateStateInvalidFieldsResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_fields")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalCatalogBookDTOToBookDTOResponse builds a value of type
// *BookDTOResponse from a value of type *catalog.BookDTO.
func marshalCatalogBookDTOToBookDTOResponse(v *catalog.BookDTO) *BookDTOResponse {
	res := &BookDTOResponse{
		ID:       v.ID,
		Title:    v.Title,
		Synopsis: v.Synopsis,
		Summary:  v.Summary,
		Price:    v.Price,
		Pages:    v.Pages,
		Isbn:     v.Isbn,
		Issue:    v.Issue,
	}
	if v.CategoryIds != nil {
		res.CategoryIds = make([]string, len(v.CategoryIds))
		for i, val := range v.CategoryIds {
			res.CategoryIds[i] = val
		}
	}
	if v.ActorIds != nil {
		res.ActorIds = make([]string, len(v.ActorIds))
		for i, val := range v.ActorIds {
			res.ActorIds[i] = val
		}
	}

	return res
}
