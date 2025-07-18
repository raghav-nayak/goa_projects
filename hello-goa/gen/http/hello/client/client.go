// Code generated by goa v3.21.1, DO NOT EDIT.
//
// hello client HTTP transport
//
// Command:
// $ goa gen hello/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the hello service endpoint HTTP clients.
type Client struct {
	// GreetHello Doer is the HTTP client used to make requests to the greetHello
	// endpoint.
	GreetHelloDoer goahttp.Doer

	// RespondToHello Doer is the HTTP client used to make requests to the
	// respondToHello endpoint.
	RespondToHelloDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the hello service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GreetHelloDoer:      doer,
		RespondToHelloDoer:  doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GreetHello returns an endpoint that makes HTTP requests to the hello service
// greetHello server.
func (c *Client) GreetHello() goa.Endpoint {
	var (
		decodeResponse = DecodeGreetHelloResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGreetHelloRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GreetHelloDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hello", "greetHello", err)
		}
		return decodeResponse(resp)
	}
}

// RespondToHello returns an endpoint that makes HTTP requests to the hello
// service respondToHello server.
func (c *Client) RespondToHello() goa.Endpoint {
	var (
		decodeResponse = DecodeRespondToHelloResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRespondToHelloRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RespondToHelloDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hello", "respondToHello", err)
		}
		return decodeResponse(resp)
	}
}
