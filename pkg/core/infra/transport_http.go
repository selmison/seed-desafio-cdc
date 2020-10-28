package infra

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	_ = EncodeErrorResponse(w, response)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(response)
}

func EncodeErrorRequest(_ context.Context, err error, w http.ResponseWriter) {
	_ = codeFrom(w, err)
}

func EncodeErrorResponse(w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok {
		return codeFrom(w, err)
	}
	return nil
}

func codeFrom(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, domain.ErrNotFound):
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case
		errors.Is(err, domain.ErrAlreadyExists),
		errors.Is(err, domain.ErrCouldNotBeEmpty),
		errors.Is(err, domain.ErrIsRequired),
		errors.Is(err, domain.ErrIsNotValidated):
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return err
}

func ToJSON(i interface{}) []byte {
	s, _ := json.Marshal(i)
	return s
}
