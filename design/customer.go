package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var CustomerDTO = Type("CustomerDTO", func() {
	Description("Customer Type")
	Attribute("id", String)
	Attribute("first_name", String)
	Attribute("last_name", String)
	Attribute("email", String, func() {
		Format(FormatEmail)
	})
	Attribute("document", String)
	Attribute("address", AddressDTO)
	Attribute("phone", String)
	Attribute("cart_ids", ArrayOf(String))
	Required("id", "first_name", "last_name", "email", "document", "address", "phone")
})

var CreateCustomerDTO = Type("CreateCustomerDTO", func() {
	Description("Customer Type")
	Attribute("first_name", String)
	Attribute("last_name", String)
	Attribute("email", String, func() {
		Format(FormatEmail)
	})
	Attribute("document", String)
	Attribute("address", AddressDTO)
	Attribute("phone", String)
	Attribute("cart_ids", ArrayOf(String))
	Required("first_name", "last_name", "email", "document", "address", "phone")
})

var AddressDTO = Type("AddressDTO", func() {
	Description("Address Type")
	Attribute("address", String)
	Attribute("complement", String)
	Attribute("city", String)
	Attribute("state_id", String)
	Attribute("cep", String)
	Required("address", "complement", "city", "state_id", "cep")
})
