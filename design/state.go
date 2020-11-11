package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var StateDTO = Type("StateDTO", func() {
	Description("State Type")
	Attribute("id", String)
	Attribute("name", String)
	Attribute("country_id")
	Required("id", "name", "country_id")
})

var StatesDTO = ArrayOf(StateDTO)

var CreateStateDTO = Type("CreateStateDTO", func() {
	Description("New State Type")
	Attribute("name", String)
	Attribute("country_id", String)
	Required("name", "country_id")
})
