// Code generated by goa v3.6.2, DO NOT EDIT.
//
// HTTP request path constructors for the healthy service.
//
// Command:
// $ goa gen goa/design -o ./

package server

// GetHealthyPath returns the URL path to the healthy service get HTTP endpoint.
func GetHealthyPath() string {
	return "/health-check"
}
