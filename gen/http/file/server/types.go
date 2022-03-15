// Code generated by goa v3.6.2, DO NOT EDIT.
//
// file HTTP server types
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	file "goa/gen/file"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UploadImageRequestBody is the type of the "file" service "upload_image"
// endpoint HTTP request body.
type UploadImageRequestBody struct {
	// file
	File []byte `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
	// file name
	FileName *string `form:"file_name,omitempty" json:"file_name,omitempty" xml:"file_name,omitempty"`
}

// UploadImageResponseBody is the type of the "file" service "upload_image"
// endpoint HTTP response body.
type UploadImageResponseBody struct {
	// image URL
	URL string `form:"url" json:"url" xml:"url"`
}

// UploadImageUnauthorizedResponseBody is the type of the "file" service
// "upload_image" endpoint HTTP response body for the "Unauthorized" error.
type UploadImageUnauthorizedResponseBody struct {
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

// NewUploadImageResponseBody builds the HTTP response body from the result of
// the "upload_image" endpoint of the "file" service.
func NewUploadImageResponseBody(res *file.ImageURL) *UploadImageResponseBody {
	body := &UploadImageResponseBody{
		URL: res.URL,
	}
	return body
}

// NewUploadImageUnauthorizedResponseBody builds the HTTP response body from
// the result of the "upload_image" endpoint of the "file" service.
func NewUploadImageUnauthorizedResponseBody(res *goa.ServiceError) *UploadImageUnauthorizedResponseBody {
	body := &UploadImageUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadImageImageFile builds a file service upload_image endpoint payload.
func NewUploadImageImageFile(body *UploadImageRequestBody) *file.ImageFile {
	v := &file.ImageFile{
		File:     body.File,
		FileName: *body.FileName,
	}

	return v
}

// ValidateUploadImageRequestBody runs the validations defined on
// upload_image_request_body
func ValidateUploadImageRequestBody(body *UploadImageRequestBody) (err error) {
	if body.File == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
	}
	if body.FileName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("file_name", "body"))
	}
	if len(body.File) > 864000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.file", body.File, len(body.File), 864000, false))
	}
	if body.FileName != nil {
		if utf8.RuneCountInString(*body.FileName) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.file_name", *body.FileName, utf8.RuneCountInString(*body.FileName), 50, false))
		}
	}
	return
}
