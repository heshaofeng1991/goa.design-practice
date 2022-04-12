// Code generated by goa v3.6.2, DO NOT EDIT.
//
// integrations HTTP server types
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	integrations "goa/gen/integrations"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AuthorizeRequestBody is the type of the "integrations" service "authorize"
// endpoint HTTP request body.
type AuthorizeRequestBody struct {
	// platform
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
	// host
	Host *string `form:"host,omitempty" json:"host,omitempty" xml:"host,omitempty"`
}

// ListResponseBody is the type of the "integrations" service "list" endpoint
// HTTP response body.
type ListResponseBody struct {
	// data
	Data []*IntegrationResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// AuthorizeResponseBody is the type of the "integrations" service "authorize"
// endpoint HTTP response body.
type AuthorizeResponseBody struct {
	// data
	Data *AuthorizePlatformRspDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// ListUnauthorizedResponseBody is the type of the "integrations" service
// "list" endpoint HTTP response body for the "Unauthorized" error.
type ListUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AuthorizeUnauthorizedResponseBody is the type of the "integrations" service
// "authorize" endpoint HTTP response body for the "Unauthorized" error.
type AuthorizeUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// IntegrationResponseBody is used to define fields on response body types.
type IntegrationResponseBody struct {
	// id
	ID *int32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// AuthorizePlatformRspDataResponseBody is used to define fields on response
// body types.
type AuthorizePlatformRspDataResponseBody struct {
	// redirect_url
	RedirectURL string `form:"redirect_url" json:"redirect_url" xml:"redirect_url"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "integrations" service.
func NewListResponseBody(res *integrations.IntegrationListResult) *ListResponseBody {
	body := &ListResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = make([]*IntegrationResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalIntegrationsIntegrationToIntegrationResponseBody(val)
		}
	}
	return body
}

// NewAuthorizeResponseBody builds the HTTP response body from the result of
// the "authorize" endpoint of the "integrations" service.
func NewAuthorizeResponseBody(res *integrations.AuthorizePlatformRsp) *AuthorizeResponseBody {
	body := &AuthorizeResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = marshalIntegrationsAuthorizePlatformRspDataToAuthorizePlatformRspDataResponseBody(res.Data)
	}
	return body
}

// NewListUnauthorizedResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "integrations" service.
func NewListUnauthorizedResponseBody(res *goa.ServiceError) *ListUnauthorizedResponseBody {
	body := &ListUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthorizeUnauthorizedResponseBody builds the HTTP response body from the
// result of the "authorize" endpoint of the "integrations" service.
func NewAuthorizeUnauthorizedResponseBody(res *goa.ServiceError) *AuthorizeUnauthorizedResponseBody {
	body := &AuthorizeUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListAuthToken builds a integrations service list endpoint payload.
func NewListAuthToken(authorization *string, token *string) *integrations.AuthToken {
	v := &integrations.AuthToken{}
	v.Authorization = authorization
	v.Token = token

	return v
}

// NewAuthorize2 builds a integrations service authorize endpoint payload.
func NewAuthorize2(body *AuthorizeRequestBody, authorization *string, token *string) *integrations.Authorize2 {
	v := &integrations.Authorize2{
		Platform: *body.Platform,
		Host:     *body.Host,
	}
	v.Authorization = authorization
	v.Token = token

	return v
}

// ValidateAuthorizeRequestBody runs the validations defined on
// AuthorizeRequestBody
func ValidateAuthorizeRequestBody(body *AuthorizeRequestBody) (err error) {
	if body.Platform == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("platform", "body"))
	}
	if body.Host == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("host", "body"))
	}
	if body.Platform != nil {
		if utf8.RuneCountInString(*body.Platform) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.platform", *body.Platform, utf8.RuneCountInString(*body.Platform), 50, false))
		}
	}
	if body.Host != nil {
		if utf8.RuneCountInString(*body.Host) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.host", *body.Host, utf8.RuneCountInString(*body.Host), 50, false))
		}
	}
	return
}