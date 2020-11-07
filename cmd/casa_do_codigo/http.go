package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
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
func handleHTTPServer(ctx context.Context, u *url.URL, catalogEndpoints *catalogGen.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

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
		catalogCreateActorHandler    *kitHttp.Server
		catalogShowActorHandler      *kitHttp.Server
		catalogCreateBookHandler     *kitHttp.Server
		catalogListBooksHandler      *kitHttp.Server
		catalogShowBookHandler       *kitHttp.Server
		catalogCreateCategoryHandler *kitHttp.Server
		catalogShowCategoryHandler   *kitHttp.Server
		catalogCreateCustomerHandler *kitHttp.Server
		catalogCreateCountryHandler  *kitHttp.Server
		catalogCreateStateHandler    *kitHttp.Server
		catalogServer                *catalogSvr.Server
	)
	{
		eh := errorHandler(logger)
		catalogCreateActorHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateActor),
			catalogKitSvr.DecodeCreateActorRequest(mux, dec),
			catalogKitSvr.EncodeCreateActorResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateActorError(enc, nil)),
		)
		catalogShowActorHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.ShowActor),
			catalogKitSvr.DecodeShowActorRequest(mux, dec),
			catalogKitSvr.EncodeShowActorResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowActorError(enc, nil)),
		)
		catalogCreateBookHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateBook),
			catalogKitSvr.DecodeCreateBookRequest(mux, dec),
			catalogKitSvr.EncodeCreateBookResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateBookError(enc, nil)),
		)
		catalogListBooksHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.ListBooks),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			catalogKitSvr.EncodeListBooksResponse(enc),
		)
		catalogShowBookHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.ShowBook),
			catalogKitSvr.DecodeShowBookRequest(mux, dec),
			catalogKitSvr.EncodeShowBookResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowBookError(enc, nil)),
		)
		catalogCreateCategoryHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateCategory),
			catalogKitSvr.DecodeCreateCategoryRequest(mux, dec),
			catalogKitSvr.EncodeCreateCategoryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCategoryError(enc, nil)),
		)
		catalogShowCategoryHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.ShowCategory),
			catalogKitSvr.DecodeShowCategoryRequest(mux, dec),
			catalogKitSvr.EncodeShowCategoryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeShowCategoryError(enc, nil)),
		)
		catalogCreateCustomerHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateCustomer),
			catalogKitSvr.DecodeCreateCustomerRequest(mux, dec),
			catalogKitSvr.EncodeCreateCustomerResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCustomerError(enc, nil)),
		)
		catalogCreateCountryHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateCountry),
			catalogKitSvr.DecodeCreateCountryRequest(mux, dec),
			catalogKitSvr.EncodeCreateCountryResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateCountryError(enc, nil)),
		)
		catalogCreateStateHandler = kitHttp.NewServer(
			endpoint.Endpoint(catalogEndpoints.CreateState),
			catalogKitSvr.DecodeCreateStateRequest(mux, dec),
			catalogKitSvr.EncodeCreateStateResponse(enc),
			kitHttp.ServerErrorEncoder(catalogKitSvr.EncodeCreateStateError(enc, nil)),
		)
		catalogServer = catalogSvr.New(catalogEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	catalogKitSvr.MountCreateActorHandler(mux, catalogCreateActorHandler)
	catalogKitSvr.MountShowActorHandler(mux, catalogShowActorHandler)
	catalogKitSvr.MountCreateBookHandler(mux, catalogCreateBookHandler)
	catalogKitSvr.MountListBooksHandler(mux, catalogListBooksHandler)
	catalogKitSvr.MountShowBookHandler(mux, catalogShowBookHandler)
	catalogKitSvr.MountCreateCategoryHandler(mux, catalogCreateCategoryHandler)
	catalogKitSvr.MountShowCategoryHandler(mux, catalogShowCategoryHandler)
	catalogKitSvr.MountCreateCustomerHandler(mux, catalogCreateCustomerHandler)
	catalogKitSvr.MountCreateCountryHandler(mux, catalogCreateCountryHandler)
	catalogKitSvr.MountCreateStateHandler(mux, catalogCreateStateHandler)

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
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
