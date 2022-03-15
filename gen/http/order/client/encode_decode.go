// Code generated by goa v3.6.2, DO NOT EDIT.
//
// order HTTP client encoders and decoders
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"bytes"
	"context"
	order "goa/gen/order"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateInboundOrderRequest instantiates a HTTP request object with
// method and path set to call the "order" service "create_inbound_order"
// endpoint
func (c *Client) BuildCreateInboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateInboundOrderOrderPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "create_inbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateInboundOrderRequest returns an encoder for requests sent to the
// order create_inbound_order server.
func EncodeCreateInboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.InboundOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "create_inbound_order", "*order.InboundOrder", v)
		}
		body := NewCreateInboundOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("order", "create_inbound_order", err)
		}
		return nil
	}
}

// DecodeCreateInboundOrderResponse returns a decoder for responses returned by
// the order create_inbound_order endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeCreateInboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeCreateInboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateInboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_inbound_order", err)
			}
			res := NewCreateInboundOrderInboundOrderRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreateInboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_inbound_order", err)
			}
			err = ValidateCreateInboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create_inbound_order", err)
			}
			return nil, NewCreateInboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "create_inbound_order", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateInboundOrderRequest instantiates a HTTP request object with
// method and path set to call the "order" service "update_inbound_order"
// endpoint
func (c *Client) BuildUpdateInboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateInboundOrderOrderPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "update_inbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateInboundOrderRequest returns an encoder for requests sent to the
// order update_inbound_order server.
func EncodeUpdateInboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.InboundOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "update_inbound_order", "*order.InboundOrder", v)
		}
		body := NewUpdateInboundOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("order", "update_inbound_order", err)
		}
		return nil
	}
}

// DecodeUpdateInboundOrderResponse returns a decoder for responses returned by
// the order update_inbound_order endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUpdateInboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateInboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UpdateInboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "update_inbound_order", err)
			}
			err = ValidateUpdateInboundOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "update_inbound_order", err)
			}
			res := NewUpdateInboundOrderUpdateResponseOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateInboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "update_inbound_order", err)
			}
			err = ValidateUpdateInboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "update_inbound_order", err)
			}
			return nil, NewUpdateInboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "update_inbound_order", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateOutboundOrderRequest instantiates a HTTP request object with
// method and path set to call the "order" service "create_outbound_order"
// endpoint
func (c *Client) BuildCreateOutboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateOutboundOrderOrderPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "create_outbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateOutboundOrderRequest returns an encoder for requests sent to the
// order create_outbound_order server.
func EncodeCreateOutboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.OutboundOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "create_outbound_order", "*order.OutboundOrder", v)
		}
		body := NewCreateOutboundOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("order", "create_outbound_order", err)
		}
		return nil
	}
}

// DecodeCreateOutboundOrderResponse returns a decoder for responses returned
// by the order create_outbound_order endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeCreateOutboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeCreateOutboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateOutboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_outbound_order", err)
			}
			res := NewCreateOutboundOrderOutboundOrderRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreateOutboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_outbound_order", err)
			}
			err = ValidateCreateOutboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create_outbound_order", err)
			}
			return nil, NewCreateOutboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "create_outbound_order", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateOutboundOrderRequest instantiates a HTTP request object with
// method and path set to call the "order" service "update_outbound_order"
// endpoint
func (c *Client) BuildUpdateOutboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateOutboundOrderOrderPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "update_outbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateOutboundOrderRequest returns an encoder for requests sent to the
// order update_outbound_order server.
func EncodeUpdateOutboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.OutboundOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "update_outbound_order", "*order.OutboundOrder", v)
		}
		body := NewUpdateOutboundOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("order", "update_outbound_order", err)
		}
		return nil
	}
}

// DecodeUpdateOutboundOrderResponse returns a decoder for responses returned
// by the order update_outbound_order endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeUpdateOutboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateOutboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UpdateOutboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "update_outbound_order", err)
			}
			err = ValidateUpdateOutboundOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "update_outbound_order", err)
			}
			res := NewUpdateOutboundOrderUpdateResponseOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateOutboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "update_outbound_order", err)
			}
			err = ValidateUpdateOutboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "update_outbound_order", err)
			}
			return nil, NewUpdateOutboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "update_outbound_order", resp.StatusCode, string(body))
		}
	}
}

// BuildCreatePickupOrderRequest instantiates a HTTP request object with method
// and path set to call the "order" service "create_pickup_order" endpoint
func (c *Client) BuildCreatePickupOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreatePickupOrderOrderPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "create_pickup_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreatePickupOrderRequest returns an encoder for requests sent to the
// order create_pickup_order server.
func EncodeCreatePickupOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.PickupOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "create_pickup_order", "*order.PickupOrder", v)
		}
		body := NewCreatePickupOrderRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("order", "create_pickup_order", err)
		}
		return nil
	}
}

// DecodeCreatePickupOrderResponse returns a decoder for responses returned by
// the order create_pickup_order endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeCreatePickupOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeCreatePickupOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreatePickupOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_pickup_order", err)
			}
			err = ValidateCreatePickupOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create_pickup_order", err)
			}
			res := NewCreatePickupOrderPickupOrderRspOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreatePickupOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create_pickup_order", err)
			}
			err = ValidateCreatePickupOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create_pickup_order", err)
			}
			return nil, NewCreatePickupOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "create_pickup_order", resp.StatusCode, string(body))
		}
	}
}

// BuildGetInboundOrderRequest instantiates a HTTP request object with method
// and path set to call the "order" service "get_inbound_order" endpoint
func (c *Client) BuildGetInboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetInboundOrderOrderPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "get_inbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetInboundOrderRequest returns an encoder for requests sent to the
// order get_inbound_order server.
func EncodeGetInboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.GetOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "get_inbound_order", "*order.GetOrder", v)
		}
		values := req.URL.Query()
		values.Add("client_order_id", p.ClientOrderID)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetInboundOrderResponse returns a decoder for responses returned by
// the order get_inbound_order endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeGetInboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeGetInboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetInboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "get_inbound_order", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateInboundOrderResponseResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "get_inbound_order", err)
			}
			res := NewGetInboundOrderInboundOrderResponseOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetInboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "get_inbound_order", err)
			}
			err = ValidateGetInboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "get_inbound_order", err)
			}
			return nil, NewGetInboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "get_inbound_order", resp.StatusCode, string(body))
		}
	}
}

// BuildGetOutboundOrderRequest instantiates a HTTP request object with method
// and path set to call the "order" service "get_outbound_order" endpoint
func (c *Client) BuildGetOutboundOrderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetOutboundOrderOrderPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "get_outbound_order", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetOutboundOrderRequest returns an encoder for requests sent to the
// order get_outbound_order server.
func EncodeGetOutboundOrderRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*order.GetOrder)
		if !ok {
			return goahttp.ErrInvalidType("order", "get_outbound_order", "*order.GetOrder", v)
		}
		values := req.URL.Query()
		values.Add("client_order_id", p.ClientOrderID)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetOutboundOrderResponse returns a decoder for responses returned by
// the order get_outbound_order endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeGetOutboundOrderResponse may return the following errors:
//	- "Unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeGetOutboundOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetOutboundOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "get_outbound_order", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateOrderRspResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "get_outbound_order", err)
			}
			res := NewGetOutboundOrderOrderRspOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetOutboundOrderUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "get_outbound_order", err)
			}
			err = ValidateGetOutboundOrderUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "get_outbound_order", err)
			}
			return nil, NewGetOutboundOrderUnauthorized(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "get_outbound_order", resp.StatusCode, string(body))
		}
	}
}

// marshalOrderItemToItemRequestBody builds a value of type *ItemRequestBody
// from a value of type *order.Item.
func marshalOrderItemToItemRequestBody(v *order.Item) *ItemRequestBody {
	res := &ItemRequestBody{
		ProductName: v.ProductName,
		ProductSku:  v.ProductSku,
		Barcode:     v.Barcode,
		Qty:         v.Qty,
	}

	return res
}

// marshalOrderShippingAddressToShippingAddressRequestBody builds a value of
// type *ShippingAddressRequestBody from a value of type *order.ShippingAddress.
func marshalOrderShippingAddressToShippingAddressRequestBody(v *order.ShippingAddress) *ShippingAddressRequestBody {
	res := &ShippingAddressRequestBody{
		FirstName:   v.FirstName,
		LastName:    v.LastName,
		PhoneNumber: v.PhoneNumber,
		CountryName: v.CountryName,
		CountryCode: v.CountryCode,
		StateName:   v.StateName,
		StateCode:   v.StateCode,
		Address1:    v.Address1,
		Address2:    v.Address2,
		CityName:    v.CityName,
		ZipCode:     v.ZipCode,
		Name:        v.Name,
	}

	return res
}

// marshalItemRequestBodyToOrderItem builds a value of type *order.Item from a
// value of type *ItemRequestBody.
func marshalItemRequestBodyToOrderItem(v *ItemRequestBody) *order.Item {
	res := &order.Item{
		ProductName: v.ProductName,
		ProductSku:  v.ProductSku,
		Barcode:     v.Barcode,
		Qty:         v.Qty,
	}

	return res
}

// marshalShippingAddressRequestBodyToOrderShippingAddress builds a value of
// type *order.ShippingAddress from a value of type *ShippingAddressRequestBody.
func marshalShippingAddressRequestBodyToOrderShippingAddress(v *ShippingAddressRequestBody) *order.ShippingAddress {
	res := &order.ShippingAddress{
		FirstName:   v.FirstName,
		LastName:    v.LastName,
		PhoneNumber: v.PhoneNumber,
		CountryName: v.CountryName,
		CountryCode: v.CountryCode,
		StateName:   v.StateName,
		StateCode:   v.StateCode,
		Address1:    v.Address1,
		Address2:    v.Address2,
		CityName:    v.CityName,
		ZipCode:     v.ZipCode,
		Name:        v.Name,
	}

	return res
}

// marshalOrderOutboundOrderItemToOutboundOrderItemRequestBody builds a value
// of type *OutboundOrderItemRequestBody from a value of type
// *order.OutboundOrderItem.
func marshalOrderOutboundOrderItemToOutboundOrderItemRequestBody(v *order.OutboundOrderItem) *OutboundOrderItemRequestBody {
	res := &OutboundOrderItemRequestBody{
		ProductName:        v.ProductName,
		ProductSku:         v.ProductSku,
		ProductPrice:       v.ProductPrice,
		Qty:                v.Qty,
		HsCode:             v.HsCode,
		DeclaredCnName:     v.DeclaredCnName,
		DeclaredEnName:     v.DeclaredEnName,
		DeclaredValueInUsd: v.DeclaredValueInUsd,
		ProductWeight:      v.ProductWeight,
		ProductLength:      v.ProductLength,
		ProductWidth:       v.ProductWidth,
		ProductHeight:      v.ProductHeight,
		Barcode:            v.Barcode,
		DeclaredValueInEur: v.DeclaredValueInEur,
	}
	if v.ProductAttributes != nil {
		res.ProductAttributes = make([]string, len(v.ProductAttributes))
		for i, val := range v.ProductAttributes {
			res.ProductAttributes[i] = val
		}
	}

	return res
}

// marshalOutboundOrderItemRequestBodyToOrderOutboundOrderItem builds a value
// of type *order.OutboundOrderItem from a value of type
// *OutboundOrderItemRequestBody.
func marshalOutboundOrderItemRequestBodyToOrderOutboundOrderItem(v *OutboundOrderItemRequestBody) *order.OutboundOrderItem {
	res := &order.OutboundOrderItem{
		ProductName:        v.ProductName,
		ProductSku:         v.ProductSku,
		ProductPrice:       v.ProductPrice,
		Qty:                v.Qty,
		HsCode:             v.HsCode,
		DeclaredCnName:     v.DeclaredCnName,
		DeclaredEnName:     v.DeclaredEnName,
		DeclaredValueInUsd: v.DeclaredValueInUsd,
		ProductWeight:      v.ProductWeight,
		ProductLength:      v.ProductLength,
		ProductWidth:       v.ProductWidth,
		ProductHeight:      v.ProductHeight,
		Barcode:            v.Barcode,
		DeclaredValueInEur: v.DeclaredValueInEur,
	}
	if v.ProductAttributes != nil {
		res.ProductAttributes = make([]string, len(v.ProductAttributes))
		for i, val := range v.ProductAttributes {
			res.ProductAttributes[i] = val
		}
	}

	return res
}

// unmarshalInboundOrderResponseResponseToOrderInboundOrderResponse builds a
// value of type *order.InboundOrderResponse from a value of type
// *InboundOrderResponseResponse.
func unmarshalInboundOrderResponseResponseToOrderInboundOrderResponse(v *InboundOrderResponseResponse) *order.InboundOrderResponse {
	res := &order.InboundOrderResponse{
		ClientOrderID:   *v.ClientOrderID,
		Status:          *v.Status,
		PlatformOrderID: *v.PlatformOrderID,
		TrackingNumber:  *v.TrackingNumber,
		TrackingURL:     *v.TrackingURL,
		Timestamp:       *v.Timestamp,
	}
	res.Items = make([]*order.Item, len(v.Items))
	for i, val := range v.Items {
		res.Items[i] = unmarshalItemResponseToOrderItem(val)
	}

	return res
}

// unmarshalItemResponseToOrderItem builds a value of type *order.Item from a
// value of type *ItemResponse.
func unmarshalItemResponseToOrderItem(v *ItemResponse) *order.Item {
	res := &order.Item{
		ProductName: *v.ProductName,
		ProductSku:  *v.ProductSku,
		Barcode:     *v.Barcode,
		Qty:         *v.Qty,
	}

	return res
}

// unmarshalOrderRspResponseToOrderOrderRsp builds a value of type
// *order.OrderRsp from a value of type *OrderRspResponse.
func unmarshalOrderRspResponseToOrderOrderRsp(v *OrderRspResponse) *order.OrderRsp {
	res := &order.OrderRsp{
		ClientOrderID:   *v.ClientOrderID,
		Status:          *v.Status,
		PlatformOrderID: *v.PlatformOrderID,
		TrackingNumber:  *v.TrackingNumber,
		TrackingURL:     *v.TrackingURL,
	}
	res.Items = make([]*order.OutboundOrderItem, len(v.Items))
	for i, val := range v.Items {
		res.Items[i] = unmarshalOutboundOrderItemResponseToOrderOutboundOrderItem(val)
	}

	return res
}

// unmarshalOutboundOrderItemResponseToOrderOutboundOrderItem builds a value of
// type *order.OutboundOrderItem from a value of type
// *OutboundOrderItemResponse.
func unmarshalOutboundOrderItemResponseToOrderOutboundOrderItem(v *OutboundOrderItemResponse) *order.OutboundOrderItem {
	res := &order.OutboundOrderItem{
		ProductName:        *v.ProductName,
		ProductSku:         *v.ProductSku,
		ProductPrice:       *v.ProductPrice,
		Qty:                *v.Qty,
		HsCode:             *v.HsCode,
		DeclaredCnName:     *v.DeclaredCnName,
		DeclaredEnName:     *v.DeclaredEnName,
		DeclaredValueInUsd: *v.DeclaredValueInUsd,
		ProductWeight:      *v.ProductWeight,
		ProductLength:      *v.ProductLength,
		ProductWidth:       *v.ProductWidth,
		ProductHeight:      *v.ProductHeight,
		Barcode:            *v.Barcode,
		DeclaredValueInEur: *v.DeclaredValueInEur,
	}
	res.ProductAttributes = make([]string, len(v.ProductAttributes))
	for i, val := range v.ProductAttributes {
		res.ProductAttributes[i] = val
	}

	return res
}
