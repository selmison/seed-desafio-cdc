package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = API("casa_do_codigo", func() {
	Title("Casa do Código Service")
	Description("Casa do Código HTTP service")
	Server("casa_do_codigo", func() {
		Description("hosts for Casa do Código Service.")
		Services("actors")
		Services("categories")
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:3333/actors")
		})
	})
})

var _ = Service("swagger", func() {
	Description("Swagger service")
	Files("/swagger.json", "../../gen/http/openapi.json")
})
