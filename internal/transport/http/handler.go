package http

import (
	"context"
	"encoding/json"
	"net/http"

	"weneedthepoh/net-worth-tracker-api/internal/endpoint"
	"weneedthepoh/net-worth-tracker-api/internal/service/health"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// errorer describes the behavior of a request or response that can contain errors
type errorer interface {
	error() error
}

// MakeHealthCheckHandler builds a go-kit http transport and returns it
func MakeHealthCheckHandler(service health.Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	e := endpoint.MakeHealthCheckEndpoint(service)

	healthHandler := kithttp.NewServer(
		e,
		decodeHealthCheckRequest,
		encodeHealthCheckResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/health", healthHandler).Methods("GET")

	r.NotFoundHandler = handleRouteNotFound(r)

	return r
}

// decodeHealthCheckRequest returns an empty healthCheck request because there are no params for this request
func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthCheckRequest{}, nil
}

// encodeHealthCheckResponse encodes any errors received from handling the request and returns
func encodeHealthCheckResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res := response.(endpoint.HealthCheckResponse)
	return json.NewEncoder(w).Encode(res)
}

// encodeError writes error headers if an error was received
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
