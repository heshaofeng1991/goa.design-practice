// Code generated by goa v3.6.2, DO NOT EDIT.
//
// wix_redirect HTTP server
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	"context"
	wixredirect "goa/gen/wix_redirect"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the wix_redirect service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	ReturnWix http.Handler
	CORS      http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the wix_redirect service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *wixredirect.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ReturnWix", "GET", "/v1/integrations/wix/return"},
			{"CORS", "OPTIONS", "/v1/integrations/wix/return"},
		},
		ReturnWix: NewReturnWixHandler(e.ReturnWix, mux, decoder, encoder, errhandler, formatter),
		CORS:      NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "wix_redirect" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ReturnWix = m(s.ReturnWix)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the wix_redirect endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountReturnWixHandler(mux, h.ReturnWix)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the wix_redirect endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountReturnWixHandler configures the mux to serve the "wix_redirect" service
// "return_wix" endpoint.
func MountReturnWixHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleWixRedirectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/integrations/wix/return", f)
}

// NewReturnWixHandler creates a HTTP handler which loads the HTTP request and
// calls the "wix_redirect" service "return_wix" endpoint.
func NewReturnWixHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeReturnWixResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "return_wix")
		ctx = context.WithValue(ctx, goa.ServiceKey, "wix_redirect")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service wix_redirect.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleWixRedirectOrigin(h)
	mux.Handle("OPTIONS", "/v1/integrations/wix/return", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleWixRedirectOrigin applies the CORS response headers corresponding to
// the origin for the service wix_redirect.
func HandleWixRedirectOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Accept, Origin, Authorization, X-Api-Version, x-nss-tenant-id")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}