package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var ActorDTO = Type("ActorDTO", func() {
	Description("Actor Type")
	Attribute("id", String)
	Attribute("name", String)
	Attribute("email", String, func() {
		Format(FormatEmail)
	})
	Attribute("description", String, func() {
		MaxLength(400)
	})
	Attribute("created_at", String)
	Required("id", "name", "email", "description", "created_at")
})
var ActorIDsDTO = ArrayOf(String)

var CreateActorDTO = Type("CreateActorDTO", func() {
	Description("New Actor Type")
	Attribute("name", String)
	Attribute("email", String, func() {
		Format(FormatEmail)
	})
	Attribute("description", String, func() {
		MaxLength(400)
	})
	Required("name", "email", "description")
})
