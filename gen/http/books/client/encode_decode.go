// Code generated by goa v3.2.4, DO NOT EDIT.
//
// books HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	books "github.com/selmison/seed-desafio-cdc/gen/books"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateBookRequest instantiates a HTTP request object with method and
// path set to call the "books" service "create_book" endpoint
func (c *Client) BuildCreateBookRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateBookBooksPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "create_book", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateBookRequest returns an encoder for requests sent to the books
// create_book server.
func EncodeCreateBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*books.CreateBookDTO)
		if !ok {
			return goahttp.ErrInvalidType("books", "create_book", "*books.CreateBookDTO", v)
		}
		body := NewCreateBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("books", "create_book", err)
		}
		return nil
	}
}

// DecodeCreateBookResponse returns a decoder for responses returned by the
// books create_book endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateBookResponse may return the following errors:
//	- "invalid_fields" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateBookResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "create_book", err)
			}
			err = ValidateCreateBookResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "create_book", err)
			}
			res := NewCreateBookBookDTOCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBookInvalidFieldsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "create_book", err)
			}
			err = ValidateCreateBookInvalidFieldsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "create_book", err)
			}
			return nil, NewCreateBookInvalidFields(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "create_book", resp.StatusCode, string(body))
		}
	}
}
