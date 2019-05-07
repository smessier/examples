// Code generated by goa v3.0.0, DO NOT EDIT.
//
// divider client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa/v3"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the divider service endpoint HTTP clients.
type Client struct {
	// IntegerDivide Doer is the HTTP client used to make requests to the
	// integer_divide endpoint.
	IntegerDivideDoer goahttp.Doer

	// Divide Doer is the HTTP client used to make requests to the divide endpoint.
	DivideDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the divider service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		IntegerDivideDoer:   doer,
		DivideDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// IntegerDivide returns an endpoint that makes HTTP requests to the divider
// service integer_divide server.
func (c *Client) IntegerDivide() goa.Endpoint {
	var (
		decodeResponse = DecodeIntegerDivideResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildIntegerDivideRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.IntegerDivideDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("divider", "integer_divide", err)
		}
		return decodeResponse(resp)
	}
}

// Divide returns an endpoint that makes HTTP requests to the divider service
// divide server.
func (c *Client) Divide() goa.Endpoint {
	var (
		decodeResponse = DecodeDivideResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDivideRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DivideDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("divider", "divide", err)
		}
		return decodeResponse(resp)
	}
}
