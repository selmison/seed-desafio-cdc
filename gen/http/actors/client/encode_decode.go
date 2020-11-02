// Code generated by goa v3.2.4, DO NOT EDIT.
//
// actors HTTP client encoders and decoders
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

	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateActorRequest instantiates a HTTP request object with method and
// path set to call the "actors" service "create_actor" endpoint
func (c *Client) BuildCreateActorRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateActorActorsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("actors", "create_actor", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateActorRequest returns an encoder for requests sent to the actors
// create_actor server.
func EncodeCreateActorRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*actors.CreateActorDTO)
		if !ok {
			return goahttp.ErrInvalidType("actors", "create_actor", "*actors.CreateActorDTO", v)
		}
		body := NewCreateActorRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("actors", "create_actor", err)
		}
		return nil
	}
}

// DecodeCreateActorResponse returns a decoder for responses returned by the
// actors create_actor endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateActorResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateActorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("actors", "create_actor", err)
			}
			err = ValidateCreateActorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("actors", "create_actor", err)
			}
			res := NewCreateActorActorDTOCreated(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("actors", "create_actor", resp.StatusCode, string(body))
		}
	}
}
