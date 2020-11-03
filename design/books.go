package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("books", func() {
	Description("The books service performs operations on books")

	Error("invalid_fields")

	Method("create_book", func() {
		Payload(CreateBookDTO)
		Result(BookDTO)
		HTTP(func() {
			POST("/books")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})
})

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
	Attribute("category_id", String)
	Attribute("actor_id", String)
	Required("id", "title", "synopsis", "price", "pages", "isbn", "issue", "category_id", "actor_id")
})

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
	Attribute("category_id", String)
	Attribute("actor_id", String)
	Required("title", "synopsis", "price", "pages", "isbn", "issue", "category_id", "actor_id")
})
