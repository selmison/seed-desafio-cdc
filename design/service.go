package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("catalog", func() {
	Description("The catalog service performs operations on catalog")

	Error("invalid_fields")
	Error("not_found")

	Method("create_actor", func() {
		Payload(CreateActorDTO)
		Result(ActorDTO)
		HTTP(func() {
			POST("/actors")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("list_actors", func() {
		Result(ActorsDTO)
		HTTP(func() {
			GET("/actors")
		})
	})

	Method("show_actor", func() {
		Payload(ShowByIDDTO)
		Result(ActorDTO)
		HTTP(func() {
			GET("/actors/{id}")
			Response("not_found", StatusNotFound)
		})
	})

	Method("create_book", func() {
		Payload(CreateBookDTO)
		Result(BookDTO)
		HTTP(func() {
			POST("/books")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("list_books", func() {
		Result(BooksDTO)
		HTTP(func() {
			GET("/books")
		})
	})

	Method("show_book", func() {
		Payload(ShowByIDDTO)
		Result(BookDTO)
		HTTP(func() {
			GET("/books/{id}")
			Response("not_found", StatusNotFound)
		})
	})

	Method("create_cart", func() {
		Payload(CreateCartDTO)
		Result(CartDTO)
		HTTP(func() {
			POST("/carts")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_category", func() {
		Payload(CreateCategoryDTO)
		Result(CategoryDTO)
		HTTP(func() {
			POST("/categories")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("list_categories", func() {
		Result(CategoriesDTO)
		HTTP(func() {
			GET("/categories")
		})
	})

	Method("show_category", func() {
		Payload(ShowByIDDTO)
		Result(CategoryDTO)
		HTTP(func() {
			GET("/categories/{id}")
			Response("not_found", StatusNotFound)
		})
	})

	Method("create_country", func() {
		Payload(CreateCountryDTO)
		Result(CountryDTO)
		HTTP(func() {
			POST("/countries")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("list_countries", func() {
		Result(CountriesDTO)
		HTTP(func() {
			GET("/countries")
		})
	})

	Method("show_country", func() {
		Payload(ShowByIDDTO)
		Result(CountryDTO)
		HTTP(func() {
			GET("/countries/{id}")
			Response("not_found", StatusNotFound)
		})
	})

	Method("create_coupon", func() {
		Payload(CreateCouponDTO)
		Result(CouponDTO)
		HTTP(func() {
			POST("/coupons")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_customer", func() {
		Payload(CreateCustomerDTO)
		Result(CustomerDTO)
		HTTP(func() {
			POST("/customers")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_purchase", func() {
		Payload(CreatePurchaseDTO)
		Result(PurchaseDTO)
		HTTP(func() {
			POST("/purchases")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_state", func() {
		Payload(CreateStateDTO)
		Result(StateDTO)
		HTTP(func() {
			POST("/states")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("list_states", func() {
		Result(StatesDTO)
		HTTP(func() {
			GET("/states")
		})
	})

	Method("show_state", func() {
		Payload(ShowByIDDTO)
		Result(StateDTO)
		HTTP(func() {
			GET("/states/{id}")
			Response("not_found", StatusNotFound)
		})
	})

})

var ShowByIDDTO = Type("ShowByIDDTO", func() {
	Attribute("id", String, "ID")
	Required("id")
})
