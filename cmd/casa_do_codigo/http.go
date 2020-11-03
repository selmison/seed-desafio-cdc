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
	kithttp "github.com/go-kit/kit/transport/http"
	actors "github.com/selmison/seed-desafio-cdc/gen/actors"
	books "github.com/selmison/seed-desafio-cdc/gen/books"
	categories "github.com/selmison/seed-desafio-cdc/gen/categories"
	actorskitsvr "github.com/selmison/seed-desafio-cdc/gen/http/actors/kitserver"
	actorssvr "github.com/selmison/seed-desafio-cdc/gen/http/actors/server"
	bookskitsvr "github.com/selmison/seed-desafio-cdc/gen/http/books/kitserver"
	bookssvr "github.com/selmison/seed-desafio-cdc/gen/http/books/server"
	categorieskitsvr "github.com/selmison/seed-desafio-cdc/gen/http/categories/kitserver"
	categoriessvr "github.com/selmison/seed-desafio-cdc/gen/http/categories/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, actorsEndpoints *actors.Endpoints, categoriesEndpoints *categories.Endpoints, booksEndpoints *books.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		actorsCreateActorHandler        *kithttp.Server
		actorsServer                    *actorssvr.Server
		categoriesCreateCategoryHandler *kithttp.Server
		categoriesServer                *categoriessvr.Server
		booksCreateBookHandler          *kithttp.Server
		booksServer                     *bookssvr.Server
	)
	{
		eh := errorHandler(logger)
		actorsCreateActorHandler = kithttp.NewServer(
			endpoint.Endpoint(actorsEndpoints.CreateActor),
			actorskitsvr.DecodeCreateActorRequest(mux, dec),
			actorskitsvr.EncodeCreateActorResponse(enc),
			kithttp.ServerErrorEncoder(actorskitsvr.EncodeCreateActorError(enc, nil)),
		)
		actorsServer = actorssvr.New(actorsEndpoints, mux, dec, enc, eh, nil)
		categoriesCreateCategoryHandler = kithttp.NewServer(
			endpoint.Endpoint(categoriesEndpoints.CreateCategory),
			categorieskitsvr.DecodeCreateCategoryRequest(mux, dec),
			categorieskitsvr.EncodeCreateCategoryResponse(enc),
			kithttp.ServerErrorEncoder(categorieskitsvr.EncodeCreateCategoryError(enc, nil)),
		)
		categoriesServer = categoriessvr.New(categoriesEndpoints, mux, dec, enc, eh, nil)
		booksCreateBookHandler = kithttp.NewServer(
			endpoint.Endpoint(booksEndpoints.CreateBook),
			bookskitsvr.DecodeCreateBookRequest(mux, dec),
			bookskitsvr.EncodeCreateBookResponse(enc),
			kithttp.ServerErrorEncoder(bookskitsvr.EncodeCreateBookError(enc, nil)),
		)
		booksServer = bookssvr.New(booksEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	actorskitsvr.MountCreateActorHandler(mux, actorsCreateActorHandler)
	categorieskitsvr.MountCreateCategoryHandler(mux, categoriesCreateCategoryHandler)
	bookskitsvr.MountCreateBookHandler(mux, booksCreateBookHandler)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range actorsServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range categoriesServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range booksServer.Mounts {
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
