package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	kitLog "github.com/go-kit/kit/log"
	kitHttp "github.com/go-kit/kit/transport/http"
	goaHttp "goa.design/goa/v3/http"
	httpMdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
	catalogKitSvr "github.com/selmison/seed-desafio-cdc/gen/http/catalog/kitserver"
	catalogSvr "github.com/selmison/seed-desafio-cdc/gen/http/catalog/server"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, catalogEndpoints *catalogGen.Endpoints, wg *sync.WaitGroup, errc chan error, logger kitLog.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goaHttp.RequestDecoder
		enc = goaHttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goaHttp.Muxer
	{
		mux = goaHttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		catalogApplyCouponHandler    *kitHttp.Server
		catalogCreateActorHandler    *kitHttp.Server
		catalogListActorsHandler     *kitHttp.Server
		catalogShowActorHandler      *kitHttp.Server
		catalogCreateBookHandler     *kitHttp.Server
		catalogListBooksHandler      *kitHttp.Server
		catalogShowBookHandler       *kitHttp.Server
		catalogCreateCartHandler     *kitHttp.Server
		catalogCreateCategoryHandler *kitHttp.Server
		catalogListCategoriesHandler *kitHttp.Server
		catalogShowCategoryHandler   *kitHttp.Server
		catalogCreateCountryHandler  *kitHttp.Server
		catalogListCountriesHandler  *kitHttp.Server
		catalogShowCountryHandler    *kitHttp.Server
		catalogCreateCouponHandler   *kitHttp.Server
		catalogCreateCustomerHandler *kitHttp.Server
		catalogCreatePurchaseHandler *kitHttp.Server
		catalogShowPurchaseHandler   *kitHttp.Server
		catalogCreateStateHandler    *kitHttp.Server
		catalogListStatesHandler     *kitHttp.Server
		catalogServer                *catalogSvr.Server
	)
	{
		eh := errorHandler(logger)
		catalogApplyCouponHandler = kitHttp.NewServer(
			catalogEndpoints.ApplyCoupon,
			catalogKitSvr.DecodeApplyCouponRequest(mux, dec),
			catalogKitSvr.EncodeApplyCouponResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeApplyCouponError(enc, nil)),
		)
		catalogCreateActorHandler = kitHttp.NewServer(
			catalogEndpoints.CreateActor,
			catalogKitSvr.DecodeCreateActorRequest(mux, dec),
			catalogKitSvr.EncodeCreateActorResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateActorError(enc, nil)),
		)
		catalogListActorsHandler = kitHttp.NewServer(
			catalogEndpoints.ListActors,
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListActorsResponse(enc),
		)
		catalogShowActorHandler = kitHttp.NewServer(
			catalogEndpoints.ShowActor,
			catalogKitSvr.DecodeShowActorRequest(mux, dec),
			catalogKitSvr.EncodeShowActorResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowActorError(enc, nil)),
		)
		catalogCreateBookHandler = kitHttp.NewServer(
			catalogEndpoints.CreateBook,
			catalogKitSvr.DecodeCreateBookRequest(mux, dec),
			catalogKitSvr.EncodeCreateBookResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateBookError(enc, nil)),
		)
		catalogListBooksHandler = kitHttp.NewServer(
			catalogEndpoints.ListBooks,
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListBooksResponse(enc),
		)
		catalogShowBookHandler = kitHttp.NewServer(
			catalogEndpoints.ShowBook,
			catalogKitSvr.DecodeShowBookRequest(mux, dec),
			catalogKitSvr.EncodeShowBookResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowBookError(enc, nil)),
		)
		catalogCreateCartHandler = kitHttp.NewServer(
			catalogEndpoints.CreateCart,
			catalogKitSvr.DecodeCreateCartRequest(mux, dec),
			catalogKitSvr.EncodeCreateCartResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCartError(enc, nil)),
		)
		catalogCreateCategoryHandler = kitHttp.NewServer(
			catalogEndpoints.CreateCategory,
			catalogKitSvr.DecodeCreateCategoryRequest(mux, dec),
			catalogKitSvr.EncodeCreateCategoryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCategoryError(enc, nil)),
		)
		catalogListCategoriesHandler = kitHttp.NewServer(
			catalogEndpoints.ListCategories,
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListCategoriesResponse(enc),
		)
		catalogShowCategoryHandler = kitHttp.NewServer(
			catalogEndpoints.ShowCategory,
			catalogKitSvr.DecodeShowCategoryRequest(mux, dec),
			catalogKitSvr.EncodeShowCategoryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowCategoryError(enc, nil)),
		)
		catalogCreateCountryHandler = kitHttp.NewServer(
			catalogEndpoints.CreateCountry,
			catalogKitSvr.DecodeCreateCountryRequest(mux, dec),
			catalogKitSvr.EncodeCreateCountryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCountryError(enc, nil)),
		)
		catalogListCountriesHandler = kitHttp.NewServer(
			catalogEndpoints.ListCountries,
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListCountriesResponse(enc),
		)
		catalogShowCountryHandler = kitHttp.NewServer(
			catalogEndpoints.ShowCountry,
			catalogKitSvr.DecodeShowCountryRequest(mux, dec),
			catalogKitSvr.EncodeShowCountryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowCountryError(enc, nil)),
		)
		catalogCreateCouponHandler = kitHttp.NewServer(
			catalogEndpoints.CreateCoupon,
			catalogKitSvr.DecodeCreateCouponRequest(mux, dec),
			catalogKitSvr.EncodeCreateCouponResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCouponError(enc, nil)),
		)
		catalogCreateCustomerHandler = kitHttp.NewServer(
			catalogEndpoints.CreateCustomer,
			catalogKitSvr.DecodeCreateCustomerRequest(mux, dec),
			catalogKitSvr.EncodeCreateCustomerResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCustomerError(enc, nil)),
		)
		catalogCreatePurchaseHandler = kitHttp.NewServer(
			catalogEndpoints.CreatePurchase,
			catalogKitSvr.DecodeCreatePurchaseRequest(mux, dec),
			catalogKitSvr.EncodeCreatePurchaseResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreatePurchaseError(enc, nil)),
		)
		catalogShowPurchaseHandler = kitHttp.NewServer(
			catalogEndpoints.ShowPurchase,
			catalogKitSvr.DecodeShowPurchaseRequest(mux, dec),
			catalogKitSvr.EncodeShowPurchaseResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowPurchaseError(enc, nil)),
		)
		catalogCreateStateHandler = kitHttp.NewServer(
			catalogEndpoints.CreateState,
			catalogKitSvr.DecodeCreateStateRequest(mux, dec),
			catalogKitSvr.EncodeCreateStateResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateStateError(enc, nil)),
		)
		catalogListStatesHandler = kitHttp.NewServer(
			catalogEndpoints.ListStates,
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListStatesResponse(enc),
		)
		catalogServer = catalogSvr.New(catalogEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	catalogKitSvr.MountApplyCouponHandler(mux, catalogApplyCouponHandler)
	catalogKitSvr.MountCreateActorHandler(mux, catalogCreateActorHandler)
	catalogKitSvr.MountListActorsHandler(mux, catalogListActorsHandler)
	catalogKitSvr.MountShowActorHandler(mux, catalogShowActorHandler)
	catalogKitSvr.MountCreateBookHandler(mux, catalogCreateBookHandler)
	catalogKitSvr.MountListBooksHandler(mux, catalogListBooksHandler)
	catalogKitSvr.MountShowBookHandler(mux, catalogShowBookHandler)
	catalogKitSvr.MountCreateCartHandler(mux, catalogCreateCartHandler)
	catalogKitSvr.MountCreateCategoryHandler(mux, catalogCreateCategoryHandler)
	catalogKitSvr.MountListCategoriesHandler(mux, catalogListCategoriesHandler)
	catalogKitSvr.MountShowCategoryHandler(mux, catalogShowCategoryHandler)
	catalogKitSvr.MountCreateCountryHandler(mux, catalogCreateCountryHandler)
	catalogKitSvr.MountListCountriesHandler(mux, catalogListCountriesHandler)
	catalogKitSvr.MountShowCountryHandler(mux, catalogShowCountryHandler)
	catalogKitSvr.MountCreateCouponHandler(mux, catalogCreateCouponHandler)
	catalogKitSvr.MountCreateCustomerHandler(mux, catalogCreateCustomerHandler)
	catalogKitSvr.MountCreatePurchaseHandler(mux, catalogCreatePurchaseHandler)
	catalogKitSvr.MountShowPurchaseHandler(mux, catalogShowPurchaseHandler)
	catalogKitSvr.MountCreateStateHandler(mux, catalogCreateStateHandler)
	catalogKitSvr.MountListStatesHandler(mux, catalogListStatesHandler)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpMdlwr.Log(logger)(handler)
		handler = httpMdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range catalogServer.Mounts {
		if err := logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)); err != nil {
			log.Printf("kit/log error: %v\n", err)
		}
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			if err := logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host)); err != nil {
				log.Printf("kit/log error: %v\n", err)
			}
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		if err := logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host)); err != nil {
			log.Printf("kit/log error: %v\n", err)
		}

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger kitLog.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		if err := logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error())); err != nil {
			log.Printf("kit/log error: %v\n", err)
		}
	}
}
