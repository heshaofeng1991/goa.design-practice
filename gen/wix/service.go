// Code generated by goa v3.6.2, DO NOT EDIT.
//
// wix service
//
// Command:
// $ goa gen goa/design -o ./

package wix

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// integrations of Wix service
type Service interface {
	// CallbackWix implements callback_wix.
	CallbackWix(context.Context, *WixCallbackArgs) (res *WixCallbackRsp, err error)
	// WebhooksProductsWix implements webhooks_products_wix.
	WebhooksProductsWix(context.Context, *WebHooksCallBackProducts) (res *WebHooksCallBackProductsResp, err error)
	// ProductsList implements products_list.
	ProductsList(context.Context, *GetProductsListReq) (res *GetProductsListResp, err error)
	// OrdersList implements orders_list.
	OrdersList(context.Context, *GetOrdersListReq) (res *GetOrdersListResp, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "wix"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"callback_wix", "webhooks_products_wix", "products_list", "orders_list"}

// GetOrdersListReq is the payload type of the wix service orders_list method.
type GetOrdersListReq struct {
	// start_date
	StartDate *string
	// end_date
	EndDate *string
}

// GetOrdersListResp is the result type of the wix service orders_list method.
type GetOrdersListResp struct {
	// success
	Success bool
}

// GetProductsListReq is the payload type of the wix service products_list
// method.
type GetProductsListReq struct {
	// start_date
	StartDate *string
	// end_date
	EndDate *string
}

// GetProductsListResp is the result type of the wix service products_list
// method.
type GetProductsListResp struct {
	// success
	Success bool
}

// WebHooksCallBackProducts is the payload type of the wix service
// webhooks_products_wix method.
type WebHooksCallBackProducts struct {
	// start_date
	StartDate *string
	// end_date
	EndDate *string
}

// WebHooksCallBackProductsResp is the result type of the wix service
// webhooks_products_wix method.
type WebHooksCallBackProductsResp struct {
	// success
	Success bool
}

// WixCallbackArgs is the payload type of the wix service callback_wix method.
type WixCallbackArgs struct {
	// code
	Code string
	// state
	State *string
	// instanceId
	InstanceID string
}

// WixCallbackRsp is the result type of the wix service callback_wix method.
type WixCallbackRsp struct {
	// success
	Success bool
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "Unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}