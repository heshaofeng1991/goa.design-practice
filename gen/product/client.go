// Code generated by goa v3.6.2, DO NOT EDIT.
//
// product client
//
// Command:
// $ goa gen goa/design -o ./

package product

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "product" service client.
type Client struct {
	BatchesCreateProductEndpoint goa.Endpoint
	UpdateProductEndpoint        goa.Endpoint
	GenerateBarcodeEndpoint      goa.Endpoint
	GenerateTokenEndpoint        goa.Endpoint
}

// NewClient initializes a "product" service client given the endpoints.
func NewClient(batchesCreateProduct, updateProduct, generateBarcode, generateToken goa.Endpoint) *Client {
	return &Client{
		BatchesCreateProductEndpoint: batchesCreateProduct,
		UpdateProductEndpoint:        updateProduct,
		GenerateBarcodeEndpoint:      generateBarcode,
		GenerateTokenEndpoint:        generateToken,
	}
}

// BatchesCreateProduct calls the "batches_create_product" endpoint of the
// "product" service.
func (c *Client) BatchesCreateProduct(ctx context.Context, p *MultiProduct) (res *MultiProductRsp, err error) {
	var ires interface{}
	ires, err = c.BatchesCreateProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MultiProductRsp), nil
}

// UpdateProduct calls the "update_product" endpoint of the "product" service.
func (c *Client) UpdateProduct(ctx context.Context, p *Product) (res *UpdateResponse, err error) {
	var ires interface{}
	ires, err = c.UpdateProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateResponse), nil
}

// GenerateBarcode calls the "generate_barcode" endpoint of the "product"
// service.
func (c *Client) GenerateBarcode(ctx context.Context, p *BarCode) (res *BarCodeRsp, err error) {
	var ires interface{}
	ires, err = c.GenerateBarcodeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BarCodeRsp), nil
}

// GenerateToken calls the "generate_token" endpoint of the "product" service.
func (c *Client) GenerateToken(ctx context.Context, p *GenerateTokenReq) (res *GenerateTokenRsp, err error) {
	var ires interface{}
	ires, err = c.GenerateTokenEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GenerateTokenRsp), nil
}
