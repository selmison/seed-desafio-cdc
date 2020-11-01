package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = API("casa-do-codigo", func() {
	Title("Service for Casa do Código")
	Description("HTTP service for Casa do Código")
	Server("cdc", func() {
		Description("cdc hosts the Casa do Codigo Service.")
		Services("actors")
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:3333/actors")
		})
	})
})

var _ = Service("actors", func() {
	Description("The actors service performs operations on actors")
	Method("create", func() {
		Payload(CreateActorDTO)
		Result(ActorDTO)
		HTTP(func() {
			POST("/actors")
			Response(StatusCreated)
		})
	})
	Files("/swagger.json", "../../gen/http/openapi.json")
})

var ActorDTO = Type("ActorDTO", func() {
	Description("Actor Type")
	Attribute("id", String)
	Attribute("name", String)
	Attribute("e-mail", String)
	Attribute("description", String, func() {
		MaxLength(400)
	})
	Attribute("created_at", String)
	Required("id", "name", "e-mail", "description", "created_at")
})

var CreateActorDTO = Type("CreateActorDTO", func() {
	Description("New Actor Type")
	Attribute("name", String)
	Attribute("e-mail", String)
	Attribute("description", String, func() {
		MaxLength(400)
	})
	Required("name", "e-mail", "description")
})
