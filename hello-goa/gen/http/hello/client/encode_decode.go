// Code generated by goa v3.21.1, DO NOT EDIT.
//
// hello HTTP client encoders and decoders
//
// Command:
// $ goa gen hello/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGreetHelloRequest instantiates a HTTP request object with method and
// path set to call the "hello" service "greetHello" endpoint
func (c *Client) BuildGreetHelloRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("hello", "greetHello", "string", v)
		}
		name = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GreetHelloHelloPath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hello", "greetHello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGreetHelloResponse returns a decoder for responses returned by the
// hello greetHello endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGreetHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hello", "greetHello", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hello", "greetHello", resp.StatusCode, string(body))
		}
	}
}

// BuildRespondToHelloRequest instantiates a HTTP request object with method
// and path set to call the "hello" service "respondToHello" endpoint
func (c *Client) BuildRespondToHelloRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("hello", "respondToHello", "string", v)
		}
		name = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RespondToHelloHelloPath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hello", "respondToHello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRespondToHelloResponse returns a decoder for responses returned by the
// hello respondToHello endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeRespondToHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hello", "respondToHello", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hello", "respondToHello", resp.StatusCode, string(body))
		}
	}
}
