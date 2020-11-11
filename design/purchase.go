package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var PurchaseDTO = Type("PurchaseDTO", func() {
	Description("Purchase Type")
	Attribute("id", String)
	Attribute("customer", CustomerDTO)
	Attribute("cart", CartDTO)
	Required("id", "customer", "cart")
})

var CreatePurchaseDTO = Type("CreatePurchaseDTO", func() {
	Description("New Purchase Type")
	Attribute("customer", CreateCustomerDTO)
	Attribute("cart", CreateCartDTO)
	Required("customer", "cart")
})
