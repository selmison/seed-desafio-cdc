package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("categories", func() {
	Description("The categories service performs operations on categories")

	Error("invalid_fields")

	Method("create_category", func() {
		Payload(CreateCategoryDTO)
		Result(CategoryDTO)
		HTTP(func() {
			POST("/categories")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})
})

var CategoryDTO = Type("CategoryDTO", func() {
	Description("Category Type")
	Attribute("id", String)
	Attribute("name", String)
	Required("id", "name")
})

var CreateCategoryDTO = Type("CreateCategoryDTO", func() {
	Description("New Category Type")
	Attribute("name", String)
	Required("name")
})
