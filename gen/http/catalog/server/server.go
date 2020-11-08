// Code generated by goa v3.2.4, DO NOT EDIT.
//
// catalog HTTP server
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	catalog "github.com/selmison/seed-desafio-cdc/gen/catalog"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the catalog service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	CreateActor    http.Handler
	ShowActor      http.Handler
	CreateBook     http.Handler
	ListBooks      http.Handler
	ShowBook       http.Handler
	CreateCategory http.Handler
	ShowCategory   http.Handler
	CreateCustomer http.Handler
	CreateCart     http.Handler
	CreateCountry  http.Handler
	CreateState    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the catalog service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *catalog.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"CreateActor", "POST", "/actors"},
			{"ShowActor", "GET", "/actors/{id}"},
			{"CreateBook", "POST", "/books"},
			{"ListBooks", "GET", "/books"},
			{"ShowBook", "GET", "/books/{id}"},
			{"CreateCategory", "POST", "/categories"},
			{"ShowCategory", "GET", "/categories/{id}"},
			{"CreateCustomer", "POST", "/customers"},
			{"CreateCart", "POST", "/carts"},
			{"CreateCountry", "POST", "/countries"},
			{"CreateState", "POST", "/states"},
		},
		CreateActor:    NewCreateActorHandler(e.CreateActor, mux, decoder, encoder, errhandler, formatter),
		ShowActor:      NewShowActorHandler(e.ShowActor, mux, decoder, encoder, errhandler, formatter),
		CreateBook:     NewCreateBookHandler(e.CreateBook, mux, decoder, encoder, errhandler, formatter),
		ListBooks:      NewListBooksHandler(e.ListBooks, mux, decoder, encoder, errhandler, formatter),
		ShowBook:       NewShowBookHandler(e.ShowBook, mux, decoder, encoder, errhandler, formatter),
		CreateCategory: NewCreateCategoryHandler(e.CreateCategory, mux, decoder, encoder, errhandler, formatter),
		ShowCategory:   NewShowCategoryHandler(e.ShowCategory, mux, decoder, encoder, errhandler, formatter),
		CreateCustomer: NewCreateCustomerHandler(e.CreateCustomer, mux, decoder, encoder, errhandler, formatter),
		CreateCart:     NewCreateCartHandler(e.CreateCart, mux, decoder, encoder, errhandler, formatter),
		CreateCountry:  NewCreateCountryHandler(e.CreateCountry, mux, decoder, encoder, errhandler, formatter),
		CreateState:    NewCreateStateHandler(e.CreateState, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "catalog" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.CreateActor = m(s.CreateActor)
	s.ShowActor = m(s.ShowActor)
	s.CreateBook = m(s.CreateBook)
	s.ListBooks = m(s.ListBooks)
	s.ShowBook = m(s.ShowBook)
	s.CreateCategory = m(s.CreateCategory)
	s.ShowCategory = m(s.ShowCategory)
	s.CreateCustomer = m(s.CreateCustomer)
	s.CreateCart = m(s.CreateCart)
	s.CreateCountry = m(s.CreateCountry)
	s.CreateState = m(s.CreateState)
}

// Mount configures the mux to serve the catalog endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateActorHandler(mux, h.CreateActor)
	MountShowActorHandler(mux, h.ShowActor)
	MountCreateBookHandler(mux, h.CreateBook)
	MountListBooksHandler(mux, h.ListBooks)
	MountShowBookHandler(mux, h.ShowBook)
	MountCreateCategoryHandler(mux, h.CreateCategory)
	MountShowCategoryHandler(mux, h.ShowCategory)
	MountCreateCustomerHandler(mux, h.CreateCustomer)
	MountCreateCartHandler(mux, h.CreateCart)
	MountCreateCountryHandler(mux, h.CreateCountry)
	MountCreateStateHandler(mux, h.CreateState)
}

// MountCreateActorHandler configures the mux to serve the "catalog" service
// "create_actor" endpoint.
func MountCreateActorHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/actors", f)
}

// NewCreateActorHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "create_actor" endpoint.
func NewCreateActorHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateActorRequest(mux, decoder)
		encodeResponse = EncodeCreateActorResponse(encoder)
		encodeError    = EncodeCreateActorError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_actor")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowActorHandler configures the mux to serve the "catalog" service
// "show_actor" endpoint.
func MountShowActorHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/actors/{id}", f)
}

// NewShowActorHandler creates a HTTP handler which loads the HTTP request and
// calls the "catalog" service "show_actor" endpoint.
func NewShowActorHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowActorRequest(mux, decoder)
		encodeResponse = EncodeShowActorResponse(encoder)
		encodeError    = EncodeShowActorError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_actor")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateBookHandler configures the mux to serve the "catalog" service
// "create_book" endpoint.
func MountCreateBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books", f)
}

// NewCreateBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "catalog" service "create_book" endpoint.
func NewCreateBookHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateBookRequest(mux, decoder)
		encodeResponse = EncodeCreateBookResponse(encoder)
		encodeError    = EncodeCreateBookError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListBooksHandler configures the mux to serve the "catalog" service
// "list_books" endpoint.
func MountListBooksHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/books", f)
}

// NewListBooksHandler creates a HTTP handler which loads the HTTP request and
// calls the "catalog" service "list_books" endpoint.
func NewListBooksHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListBooksResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_books")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowBookHandler configures the mux to serve the "catalog" service
// "show_book" endpoint.
func MountShowBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/books/{id}", f)
}

// NewShowBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "catalog" service "show_book" endpoint.
func NewShowBookHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowBookRequest(mux, decoder)
		encodeResponse = EncodeShowBookResponse(encoder)
		encodeError    = EncodeShowBookError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateCategoryHandler configures the mux to serve the "catalog" service
// "create_category" endpoint.
func MountCreateCategoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/categories", f)
}

// NewCreateCategoryHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "create_category" endpoint.
func NewCreateCategoryHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateCategoryRequest(mux, decoder)
		encodeResponse = EncodeCreateCategoryResponse(encoder)
		encodeError    = EncodeCreateCategoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_category")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowCategoryHandler configures the mux to serve the "catalog" service
// "show_category" endpoint.
func MountShowCategoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/categories/{id}", f)
}

// NewShowCategoryHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "show_category" endpoint.
func NewShowCategoryHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowCategoryRequest(mux, decoder)
		encodeResponse = EncodeShowCategoryResponse(encoder)
		encodeError    = EncodeShowCategoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_category")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateCustomerHandler configures the mux to serve the "catalog" service
// "create_customer" endpoint.
func MountCreateCustomerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/customers", f)
}

// NewCreateCustomerHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "create_customer" endpoint.
func NewCreateCustomerHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateCustomerRequest(mux, decoder)
		encodeResponse = EncodeCreateCustomerResponse(encoder)
		encodeError    = EncodeCreateCustomerError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_customer")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateCartHandler configures the mux to serve the "catalog" service
// "create_cart" endpoint.
func MountCreateCartHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/carts", f)
}

// NewCreateCartHandler creates a HTTP handler which loads the HTTP request and
// calls the "catalog" service "create_cart" endpoint.
func NewCreateCartHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateCartRequest(mux, decoder)
		encodeResponse = EncodeCreateCartResponse(encoder)
		encodeError    = EncodeCreateCartError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_cart")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateCountryHandler configures the mux to serve the "catalog" service
// "create_country" endpoint.
func MountCreateCountryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/countries", f)
}

// NewCreateCountryHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "create_country" endpoint.
func NewCreateCountryHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateCountryRequest(mux, decoder)
		encodeResponse = EncodeCreateCountryResponse(encoder)
		encodeError    = EncodeCreateCountryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_country")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateStateHandler configures the mux to serve the "catalog" service
// "create_state" endpoint.
func MountCreateStateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/states", f)
}

// NewCreateStateHandler creates a HTTP handler which loads the HTTP request
// and calls the "catalog" service "create_state" endpoint.
func NewCreateStateHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateStateRequest(mux, decoder)
		encodeResponse = EncodeCreateStateResponse(encoder)
		encodeError    = EncodeCreateStateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_state")
		ctx = context.WithValue(ctx, goa.ServiceKey, "catalog")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
