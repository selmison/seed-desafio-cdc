package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var CartDTO = Type("CartDTO", func() {
	Description("Cart Type")
	Attribute("id", String)
	Attribute("total", Float32)
	Attribute("items", ItemsDTO)
	Attribute("customer_id", String)
	Required("id", "total", "items", "customer_id")
})

var CreateCartDTO = Type("CreateCartDTO", func() {
	Description("Cart Type")
	Attribute("total", Float32)
	Attribute("items", ItemsDTO)
	Attribute("customer_id", String)
	Required("total", "items", "customer_id")
})

var ItemsDTO = ArrayOf(ItemDTO)

var ItemDTO = Type("ItemDTO", func() {
	Description("Item Type")
	Attribute("book_id", String)
	Attribute("amount", Int32)
	Required("book_id", "amount")
})
