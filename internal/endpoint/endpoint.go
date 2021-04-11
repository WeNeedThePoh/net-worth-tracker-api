package endpoint

import (
	"context"

	"weneedthepoh/net-worth-tracker-api/internal/service/health"

	"github.com/go-kit/kit/endpoint"
)

// HealthCheckRequest has no parameters, but we still generate an empty struct to represent it
type HealthCheckRequest struct{}

// HealthCheckResponse represents an HTTP response from the health endpoint
type HealthCheckResponse struct {
	Database string `json:"database"`
}

// HealthCheckErrorResponse represents an HTTP error response from the health endpoint containing any errors
type HealthCheckErrorResponse struct {
	Error error `json:"error"`
}

// error is an implementation of the errorer interface allowing us to encode errors received from the service
func (r HealthCheckErrorResponse) error() error { return r.Error }

// MakeHealthCheckEndpoint returns a go-kit endpoint, wrapping the health response
func MakeHealthCheckEndpoint(service health.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		dbStatus := service.CheckHealth()
		return HealthCheckResponse{
			Database: dbStatus,
		}, nil
	}
}
