// Code generated by goa v3.6.2, DO NOT EDIT.
//
// platform_product HTTP client types
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	platformproduct "goa/gen/platform_product"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// ConvertRequestBody is the type of the "platform_product" service "convert"
// endpoint HTTP request body.
type ConvertRequestBody struct {
	// platform_product_ids
	PlatformProductIds []int32 `form:"platform_product_ids" json:"platform_product_ids" xml:"platform_product_ids"`
}

// MappingsRequestBody is the type of the "platform_product" service "mappings"
// endpoint HTTP request body.
type MappingsRequestBody struct {
	// platform_product_ids
	PlatformProductIds []int32 `form:"platform_product_ids" json:"platform_product_ids" xml:"platform_product_ids"`
	// products
	Products []*UpdateMappingsProductRequestBody `form:"products" json:"products" xml:"products"`
}

// PlatformProductResponseBody is the type of the "platform_product" service
// "platform_product" endpoint HTTP response body.
type PlatformProductResponseBody struct {
	// data
	Data *MultiPlatformProductResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code *int `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ConvertResponseBody is the type of the "platform_product" service "convert"
// endpoint HTTP response body.
type ConvertResponseBody struct {
	// data
	Data []*ConvertPlatformProductsResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code *int `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// MappingsResponseBody is the type of the "platform_product" service
// "mappings" endpoint HTTP response body.
type MappingsResponseBody struct {
	// data
	Data interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code *int `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PlatformProductUnauthorizedResponseBody is the type of the
// "platform_product" service "platform_product" endpoint HTTP response body
// for the "Unauthorized" error.
type PlatformProductUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ConvertUnauthorizedResponseBody is the type of the "platform_product"
// service "convert" endpoint HTTP response body for the "Unauthorized" error.
type ConvertUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// MappingsUnauthorizedResponseBody is the type of the "platform_product"
// service "mappings" endpoint HTTP response body for the "Unauthorized" error.
type MappingsUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// MultiPlatformProductResponseBody is used to define fields on response body
// types.
type MultiPlatformProductResponseBody struct {
	// MetaData info
	Meta *MetaDataResponseBody `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// product info
	List []*ListingResponseBody `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
}

// MetaDataResponseBody is used to define fields on response body types.
type MetaDataResponseBody struct {
	// current
	Current *int `form:"current,omitempty" json:"current,omitempty" xml:"current,omitempty"`
	// page_size
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// total
	Total *int `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
}

// ListingResponseBody is used to define fields on response body types.
type ListingResponseBody struct {
	// listing sku
	ListingSku *string `form:"listing_sku,omitempty" json:"listing_sku,omitempty" xml:"listing_sku,omitempty"`
	// barcode
	Barcode *string `form:"barcode,omitempty" json:"barcode,omitempty" xml:"barcode,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// images
	Images []string `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
	// vendor
	Vendor *string `form:"vendor,omitempty" json:"vendor,omitempty" xml:"vendor,omitempty"`
	// platform status
	PlatformStatus *int `form:"platform_status,omitempty" json:"platform_status,omitempty" xml:"platform_status,omitempty"`
	// id
	ID *int32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// store
	Store *StoreResponseBody `form:"store,omitempty" json:"store,omitempty" xml:"store,omitempty"`
	// mappings
	Mappings []*MappingResponseBody `form:"mappings,omitempty" json:"mappings,omitempty" xml:"mappings,omitempty"`
	// is_mapping
	IsMapping *bool `form:"is_mapping,omitempty" json:"is_mapping,omitempty" xml:"is_mapping,omitempty"`
}

// StoreResponseBody is used to define fields on response body types.
type StoreResponseBody struct {
	// id
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// platform
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
}

// MappingResponseBody is used to define fields on response body types.
type MappingResponseBody struct {
	// id
	ID *int32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// platform_product_id
	PlatformProductID *int32 `form:"platform_product_id,omitempty" json:"platform_product_id,omitempty" xml:"platform_product_id,omitempty"`
	// product_id
	ProductID *int32 `form:"product_id,omitempty" json:"product_id,omitempty" xml:"product_id,omitempty"`
	// product_sku
	ProductSku *string `form:"product_sku,omitempty" json:"product_sku,omitempty" xml:"product_sku,omitempty"`
	// qty
	Qty *int32 `form:"qty,omitempty" json:"qty,omitempty" xml:"qty,omitempty"`
	// product_name
	ProductName *string `form:"product_name,omitempty" json:"product_name,omitempty" xml:"product_name,omitempty"`
	// images
	Images []string `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
}

// ConvertPlatformProductsResDataResponseBody is used to define fields on
// response body types.
type ConvertPlatformProductsResDataResponseBody struct {
	// id
	ID *int32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// success
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateMappingsProductRequestBody is used to define fields on request body
// types.
type UpdateMappingsProductRequestBody struct {
	// product_id
	ProductID int32 `form:"product_id" json:"product_id" xml:"product_id"`
	// qty
	Qty int32 `form:"qty" json:"qty" xml:"qty"`
}

// NewConvertRequestBody builds the HTTP request body from the payload of the
// "convert" endpoint of the "platform_product" service.
func NewConvertRequestBody(p *platformproduct.ConvertPlatformProductsReq) *ConvertRequestBody {
	body := &ConvertRequestBody{}
	if p.PlatformProductIds != nil {
		body.PlatformProductIds = make([]int32, len(p.PlatformProductIds))
		for i, val := range p.PlatformProductIds {
			body.PlatformProductIds[i] = val
		}
	}
	return body
}

// NewMappingsRequestBody builds the HTTP request body from the payload of the
// "mappings" endpoint of the "platform_product" service.
func NewMappingsRequestBody(p *platformproduct.UpdateMappingsReq) *MappingsRequestBody {
	body := &MappingsRequestBody{}
	if p.PlatformProductIds != nil {
		body.PlatformProductIds = make([]int32, len(p.PlatformProductIds))
		for i, val := range p.PlatformProductIds {
			body.PlatformProductIds[i] = val
		}
	}
	if p.Products != nil {
		body.Products = make([]*UpdateMappingsProductRequestBody, len(p.Products))
		for i, val := range p.Products {
			body.Products[i] = marshalPlatformproductUpdateMappingsProductToUpdateMappingsProductRequestBody(val)
		}
	}
	return body
}

// NewPlatformProductMultiPlatformProductRspOK builds a "platform_product"
// service "platform_product" endpoint result from a HTTP "OK" response.
func NewPlatformProductMultiPlatformProductRspOK(body *PlatformProductResponseBody) *platformproduct.MultiPlatformProductRsp {
	v := &platformproduct.MultiPlatformProductRsp{
		Code:    *body.Code,
		Message: *body.Message,
	}
	if body.Data != nil {
		v.Data = unmarshalMultiPlatformProductResponseBodyToPlatformproductMultiPlatformProduct(body.Data)
	}

	return v
}

// NewPlatformProductUnauthorized builds a platform_product service
// platform_product endpoint Unauthorized error.
func NewPlatformProductUnauthorized(body *PlatformProductUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewConvertPlatformProductsResOK builds a "platform_product" service
// "convert" endpoint result from a HTTP "OK" response.
func NewConvertPlatformProductsResOK(body *ConvertResponseBody) *platformproduct.ConvertPlatformProductsRes {
	v := &platformproduct.ConvertPlatformProductsRes{
		Code:    *body.Code,
		Message: *body.Message,
	}
	if body.Data != nil {
		v.Data = make([]*platformproduct.ConvertPlatformProductsResData, len(body.Data))
		for i, val := range body.Data {
			v.Data[i] = unmarshalConvertPlatformProductsResDataResponseBodyToPlatformproductConvertPlatformProductsResData(val)
		}
	}

	return v
}

// NewConvertUnauthorized builds a platform_product service convert endpoint
// Unauthorized error.
func NewConvertUnauthorized(body *ConvertUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewMappingsUpdateMappingsResOK builds a "platform_product" service
// "mappings" endpoint result from a HTTP "OK" response.
func NewMappingsUpdateMappingsResOK(body *MappingsResponseBody) *platformproduct.UpdateMappingsRes {
	v := &platformproduct.UpdateMappingsRes{
		Data:    body.Data,
		Code:    *body.Code,
		Message: *body.Message,
	}

	return v
}

// NewMappingsUnauthorized builds a platform_product service mappings endpoint
// Unauthorized error.
func NewMappingsUnauthorized(body *MappingsUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidatePlatformProductResponseBody runs the validations defined on
// platform_product_response_body
func ValidatePlatformProductResponseBody(body *PlatformProductResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Data != nil {
		if err2 := ValidateMultiPlatformProductResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateConvertResponseBody runs the validations defined on
// ConvertResponseBody
func ValidateConvertResponseBody(body *ConvertResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateConvertPlatformProductsResDataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateMappingsResponseBody runs the validations defined on
// MappingsResponseBody
func ValidateMappingsResponseBody(body *MappingsResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePlatformProductUnauthorizedResponseBody runs the validations defined
// on platform_product_Unauthorized_response_body
func ValidatePlatformProductUnauthorizedResponseBody(body *PlatformProductUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateConvertUnauthorizedResponseBody runs the validations defined on
// convert_Unauthorized_response_body
func ValidateConvertUnauthorizedResponseBody(body *ConvertUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateMappingsUnauthorizedResponseBody runs the validations defined on
// mappings_Unauthorized_response_body
func ValidateMappingsUnauthorizedResponseBody(body *MappingsUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateMultiPlatformProductResponseBody runs the validations defined on
// MultiPlatformProductResponseBody
func ValidateMultiPlatformProductResponseBody(body *MultiPlatformProductResponseBody) (err error) {
	if body.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
	}
	if body.List == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("list", "body"))
	}
	if body.Meta != nil {
		if err2 := ValidateMetaDataResponseBody(body.Meta); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.List {
		if e != nil {
			if err2 := ValidateListingResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateMetaDataResponseBody runs the validations defined on
// MetaDataResponseBody
func ValidateMetaDataResponseBody(body *MetaDataResponseBody) (err error) {
	if body.Current == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("current", "body"))
	}
	if body.PageSize == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page_size", "body"))
	}
	if body.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "body"))
	}
	if body.Current != nil {
		if *body.Current < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.current", *body.Current, 1, true))
		}
	}
	if body.PageSize != nil {
		if *body.PageSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.page_size", *body.PageSize, 1, true))
		}
	}
	return
}

// ValidateListingResponseBody runs the validations defined on
// ListingResponseBody
func ValidateListingResponseBody(body *ListingResponseBody) (err error) {
	if body.ListingSku != nil {
		if utf8.RuneCountInString(*body.ListingSku) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.listing_sku", *body.ListingSku, utf8.RuneCountInString(*body.ListingSku), 50, false))
		}
	}
	if body.Barcode != nil {
		if utf8.RuneCountInString(*body.Barcode) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.barcode", *body.Barcode, utf8.RuneCountInString(*body.Barcode), 50, false))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	if body.Store != nil {
		if err2 := ValidateStoreResponseBody(body.Store); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.Mappings {
		if e != nil {
			if err2 := ValidateMappingResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoreResponseBody runs the validations defined on StoreResponseBody
func ValidateStoreResponseBody(body *StoreResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Platform == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("platform", "body"))
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	return
}

// ValidateMappingResponseBody runs the validations defined on
// MappingResponseBody
func ValidateMappingResponseBody(body *MappingResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.PlatformProductID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("platform_product_id", "body"))
	}
	if body.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("product_id", "body"))
	}
	if body.ProductSku == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("product_sku", "body"))
	}
	if body.Qty == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("qty", "body"))
	}
	if body.ProductName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("product_name", "body"))
	}
	if body.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "body"))
	}
	return
}

// ValidateConvertPlatformProductsResDataResponseBody runs the validations
// defined on ConvertPlatformProductsResDataResponseBody
func ValidateConvertPlatformProductsResDataResponseBody(body *ConvertPlatformProductsResDataResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}