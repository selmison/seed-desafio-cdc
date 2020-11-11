package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var CategoryDTO = Type("CategoryDTO", func() {
	Description("Category Type")
	Attribute("id", String)
	Attribute("name", String)
	Required("id", "name")
})

var CategoriesDTO = ArrayOf(CategoryDTO)

var CreateCategoryDTO = Type("CreateCategoryDTO", func() {
	Description("New Category Type")
	Attribute("name", String)
	Required("name")
})
