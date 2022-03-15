package design

import . "goa.design/goa/v3/dsl" //nolint:revive

var _ = Service("healthy", func() {
	Description("The healthy check service")

	HTTP(func() {
		Path("/health-check")
	})

	Method("get", func() {
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
})
