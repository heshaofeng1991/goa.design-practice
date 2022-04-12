// Code generated by goa v3.6.2, DO NOT EDIT.
//
// auth client HTTP transport
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the auth service endpoint HTTP clients.
type Client struct {
	// GenerateToken Doer is the HTTP client used to make requests to the
	// generate_token endpoint.
	GenerateTokenDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the auth service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GenerateTokenDoer:   doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GenerateToken returns an endpoint that makes HTTP requests to the auth
// service generate_token server.
func (c *Client) GenerateToken() goa.Endpoint {
	var (
		encodeRequest  = EncodeGenerateTokenRequest(c.encoder)
		decodeResponse = DecodeGenerateTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGenerateTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GenerateTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("auth", "generate_token", err)
		}
		return decodeResponse(resp)
	}
}