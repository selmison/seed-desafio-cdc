// Code generated by goa v3.2.4, DO NOT EDIT.
//
// categories HTTP client encoders and decoders
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

	categories "github.com/selmison/seed-desafio-cdc/gen/categories"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateCategoryRequest instantiates a HTTP request object with method
// and path set to call the "categories" service "create_category" endpoint
func (c *Client) BuildCreateCategoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateCategoryCategoriesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("categories", "create_category", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateCategoryRequest returns an encoder for requests sent to the
// categories create_category server.
func EncodeCreateCategoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*categories.CreateCategoryDTO)
		if !ok {
			return goahttp.ErrInvalidType("categories", "create_category", "*categories.CreateCategoryDTO", v)
		}
		body := NewCreateCategoryRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("categories", "create_category", err)
		}
		return nil
	}
}

// DecodeCreateCategoryResponse returns a decoder for responses returned by the
// categories create_category endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeCreateCategoryResponse may return the following errors:
//	- "invalid_fields" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateCategoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateCategoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("categories", "create_category", err)
			}
			err = ValidateCreateCategoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("categories", "create_category", err)
			}
			res := NewCreateCategoryCategoryDTOCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateCategoryInvalidFieldsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("categories", "create_category", err)
			}
			err = ValidateCreateCategoryInvalidFieldsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("categories", "create_category", err)
			}
			return nil, NewCreateCategoryInvalidFields(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("categories", "create_category", resp.StatusCode, string(body))
		}
	}
}