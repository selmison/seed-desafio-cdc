// Code generated by goa v3.2.4, DO NOT EDIT.
//
// books client HTTP transport
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the books service endpoint HTTP clients.
type Client struct {
	// CreateBook Doer is the HTTP client used to make requests to the create_book
	// endpoint.
	CreateBookDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the books service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateBookDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreateBook returns an endpoint that makes HTTP requests to the books service
// create_book server.
func (c *Client) CreateBook() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeCreateBookRequest(c.encoder)
		decodeResponse = DecodeCreateBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "create_book", err)
		}
		return decodeResponse(resp)
	}
}
