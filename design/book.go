package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var BookDTO = Type("BookDTO", func() {
	Description("Book Type")
	Attribute("id", String)
	Attribute("title", String)
	Attribute("synopsis", String, func() {
		MaxLength(500)
	})
	Attribute("summary", String)
	Attribute("price", Float32, func() {
		Minimum(20)
	})
	Attribute("pages", Int, func() {
		Minimum(100)
	})
	Attribute("isbn", String)
	Attribute("issue", String)
	Attribute("category_ids", ArrayOf(String))
	Attribute("actor_ids", ArrayOf(String))
	Required("id", "title", "synopsis", "price", "pages", "isbn", "issue", "category_ids", "actor_ids")
})

var BooksDTO = ArrayOf(BookDTO)

var CreateBookDTO = Type("CreateBookDTO", func() {
	Description("New Book Type")
	Attribute("title", String)
	Attribute("synopsis", String, func() {
		MaxLength(500)
	})
	Attribute("summary", String)
	Attribute("price", Float32, func() {
		Minimum(20)
	})
	Attribute("pages", Int, func() {
		Minimum(100)
	})
	Attribute("isbn", String)
	Attribute("issue", String)
	Attribute("category_ids", ArrayOf(String))
	Attribute("actor_ids", ArrayOf(String))
	Required("title", "synopsis", "price", "pages", "isbn", "issue", "category_ids", "actor_ids")
})
