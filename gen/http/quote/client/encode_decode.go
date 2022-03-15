// Code generated by goa v3.6.2, DO NOT EDIT.
//
// quote HTTP client encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"bytes"
	"context"
	"fmt"
	quote "goa/gen/quote"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "quote" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetQuotePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("quote", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the quote get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*quote.GetQuote)
		if !ok {
			return goahttp.ErrInvalidType("quote", "get", "*quote.GetQuote", v)
		}
		values := req.URL.Query()
		values.Add("origin_country", p.OriginCountry)
		values.Add("dest_country", p.DestCountry)
		values.Add("dest_state", p.DestState)
		values.Add("dest_zip_code", p.DestZipCode)
		values.Add("weight", fmt.Sprintf("%v", p.Weight))
		values.Add("length", fmt.Sprintf("%v", p.Length))
		values.Add("width", fmt.Sprintf("%v", p.Width))
		values.Add("height", fmt.Sprintf("%v", p.Height))
		for _, value := range p.ProductAttributes {
			values.Add("product_attributes", value)
		}
		if p.Factory != nil {
			values.Add("factory", *p.Factory)
		}
		if p.Date != nil {
			values.Add("date", *p.Date)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the quote get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("quote", "get", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateQuoteResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("quote", "get", err)
			}
			res := NewGetQuoteOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("quote", "get", err)
			}
			err = ValidateGetUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("quote", "get", err)
			}
			return nil, NewGetUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("quote", "get", resp.StatusCode, string(body))
		}
	}
}

// unmarshalQuoteResponseToQuoteQuote builds a value of type *quote.Quote from
// a value of type *QuoteResponse.
func unmarshalQuoteResponseToQuoteQuote(v *QuoteResponse) *quote.Quote {
	res := &quote.Quote{
		ChannelName:   *v.ChannelName,
		ChannelID:     *v.ChannelID,
		Type:          *v.Type,
		MinNormalDays: *v.MinNormalDays,
		MaxNormalDays: *v.MaxNormalDays,
		TotalCost:     *v.TotalCost,
		Currency:      *v.Currency,
		Weight:        *v.Weight,
	}

	return res
}
