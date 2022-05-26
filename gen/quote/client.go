// Code generated by goa v3.6.2, DO NOT EDIT.
//
// quote client
//
// Command:
// $ goa gen goa/design -o ./

package quote

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "quote" service client.
type Client struct {
	UpdateChannelCostStatusEndpoint goa.Endpoint
}

// NewClient initializes a "quote" service client given the endpoints.
func NewClient(updateChannelCostStatus goa.Endpoint) *Client {
	return &Client{
		UpdateChannelCostStatusEndpoint: updateChannelCostStatus,
	}
}

// UpdateChannelCostStatus calls the "UpdateChannelCostStatus" endpoint of the
// "quote" service.
func (c *Client) UpdateChannelCostStatus(ctx context.Context, p *UpdateChannelCostStatusReq) (res *UpdateChannelCostStatusRsp, err error) {
	var ires interface{}
	ires, err = c.UpdateChannelCostStatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateChannelCostStatusRsp), nil
}
