package infra

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/service"
	"github.com/selmison/seed-desafio-cdc/pkg/core/infra"
)

func NewActorRoutes(router infra.Router, svc service.Service, logger kitLog.Logger) {
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(infra.EncodeErrorRequest),
	}
	router.AddRoute("GET", "/actors",
		httpTransport.NewServer(
			LoggingMiddleware(kitLog.With(logger, "method", "List"))(MakeListEndpoint(svc)),
			decodeActorListRequest,
			infra.EncodeResponse,
			options...,
		),
	)
	router.AddRoute("GET", "/actors/:id",
		httpTransport.NewServer(
			LoggingMiddleware(kitLog.With(logger, "method", "Show"))(MakeShowEndpoint(svc)),
			decodeActorShowRequest,
			encodeActorShowResponse,
			options...,
		),
	)
	router.AddRoute("POST", "/actors/",
		httpTransport.NewServer(
			LoggingMiddleware(kitLog.With(logger, "method", "Create"))(MakeCreateEndpoint(svc)),
			decodeActorCreateRequest,
			encodeActorCreateResponse,
			options...,
		),
	)
	router.AddRoute("PUT", "/actors/:id",
		httpTransport.NewServer(
			LoggingMiddleware(kitLog.With(logger, "method", "Update"))(MakeUpdateEndpoint(svc)),
			decodeActorUpdateRequest,
			encodeActorUpdateResponse,
			options...,
		),
	)
	router.AddRoute("DELETE", "/actors/:id",
		httpTransport.NewServer(
			LoggingMiddleware(kitLog.With(logger, "method", "Destroy"))(MakeDestroyEndpoint(svc)),
			decodeActorDestroyRequest,
			encodeActorDestroyResponse,
			options...,
		),
	)
}

func decodeActorCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println()
		}
	}()
	var request CreateRequest
	body := r.Body
	bs, _ := ioutil.ReadAll(body)
	if err := json.Unmarshal(bs, &request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeActorCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	err := infra.EncodeErrorResponse(w, response)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeActorDestroyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeActorDestroyResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	_ = infra.EncodeErrorResponse(w, response)
	return nil
}

func decodeActorListRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeActorShowRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeActorShowResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	showResponse, _ := response.(ShowResponse)
	if err := infra.EncodeErrorResponse(w, showResponse.Err); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeActorUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println()
		}
	}()
	var request UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeActorUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	_ = infra.EncodeErrorResponse(w, response)
	return nil
}
