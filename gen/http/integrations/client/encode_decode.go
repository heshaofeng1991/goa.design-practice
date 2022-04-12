// Code generated by goa v3.6.2, DO NOT EDIT.
//
// integrations HTTP client encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"bytes"
	"context"
	integrations "goa/gen/integrations"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "integrations" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListIntegrationsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("integrations", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the integrations
// list server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*integrations.AuthToken)
		if !ok {
			return goahttp.ErrInvalidType("integrations", "list", "*integrations.AuthToken", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the
// integrations list endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("integrations", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("integrations", "list", err)
			}
			res := NewListIntegrationListResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("integrations", "list", err)
			}
			err = ValidateListUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("integrations", "list", err)
			}
			return nil, NewListUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("integrations", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildAuthorizeRequest instantiates a HTTP request object with method and
// path set to call the "integrations" service "authorize" endpoint
func (c *Client) BuildAuthorizeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthorizeIntegrationsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("integrations", "authorize", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthorizeRequest returns an encoder for requests sent to the
// integrations authorize server.
func EncodeAuthorizeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*integrations.Authorize2)
		if !ok {
			return goahttp.ErrInvalidType("integrations", "authorize", "*integrations.Authorize2", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewAuthorizeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("integrations", "authorize", err)
		}
		return nil
	}
}

// DecodeAuthorizeResponse returns a decoder for responses returned by the
// integrations authorize endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeAuthorizeResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeAuthorizeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body AuthorizeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("integrations", "authorize", err)
			}
			err = ValidateAuthorizeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("integrations", "authorize", err)
			}
			res := NewAuthorizePlatformRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body AuthorizeUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("integrations", "authorize", err)
			}
			err = ValidateAuthorizeUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("integrations", "authorize", err)
			}
			return nil, NewAuthorizeUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("integrations", "authorize", resp.StatusCode, string(body))
		}
	}
}

// unmarshalIntegrationResponseBodyToIntegrationsIntegration builds a value of
// type *integrations.Integration from a value of type *IntegrationResponseBody.
func unmarshalIntegrationResponseBodyToIntegrationsIntegration(v *IntegrationResponseBody) *integrations.Integration {
	if v == nil {
		return nil
	}
	res := &integrations.Integration{
		ID:     v.ID,
		Name:   v.Name,
		Status: v.Status,
	}

	return res
}

// unmarshalAuthorizePlatformRspDataResponseBodyToIntegrationsAuthorizePlatformRspData
// builds a value of type *integrations.AuthorizePlatformRspData from a value
// of type *AuthorizePlatformRspDataResponseBody.
func unmarshalAuthorizePlatformRspDataResponseBodyToIntegrationsAuthorizePlatformRspData(v *AuthorizePlatformRspDataResponseBody) *integrations.AuthorizePlatformRspData {
	if v == nil {
		return nil
	}
	res := &integrations.AuthorizePlatformRspData{
		RedirectURL: *v.RedirectURL,
	}

	return res
}