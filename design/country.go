package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var CountryDTO = Type("CountryDTO", func() {
	Description("Country Type")
	Attribute("id", String)
	Attribute("name", String)
	Attribute("state_ids", String)
	Required("id", "name")
})

var CountriesDTO = ArrayOf(CountryDTO)

var CreateCountryDTO = Type("CreateCountryDTO", func() {
	Description("New Country Type")
	Attribute("name", String)
	Required("name")
})
