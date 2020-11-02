package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("categories", func() {
	Description("The categories service performs operations on categories")
	Method("create_category", func() {
		Payload(CreateCategoryDTO)
		Result(CategoryDTO)
		HTTP(func() {
			POST("/categories")
			Response(StatusCreated)
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
