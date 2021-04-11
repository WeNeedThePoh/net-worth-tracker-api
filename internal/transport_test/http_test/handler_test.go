package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"weneedthepoh/net-worth-tracker-api/internal/endpoint"
	"weneedthepoh/net-worth-tracker-api/internal/service/health"
	kithttp "weneedthepoh/net-worth-tracker-api/internal/transport/http"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestMakeHealthCheckHandler(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Mock DB pool connection
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	s := health.NewService(db)

	req := httptest.NewRequest("GET", "/health", nil)
	r := httptest.NewRecorder()

	expected := endpoint.HealthCheckResponse{
		Database: "OK",
	}

	// when
	h := kithttp.MakeHealthCheckHandler(s, logger)
	h.ServeHTTP(r, req)

	var res endpoint.HealthCheckResponse
	err := json.NewDecoder(r.Body).Decode(&res)

	// then
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.Result().StatusCode)
	assert.Equal(t, expected, res)
}

func TestMakeHealthCheckHandlerOnRouteNotFound(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Mock DB pool connection
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	s := health.NewService(db)

	req := httptest.NewRequest("GET", "/health/not/found", nil)
	r := httptest.NewRecorder()

	// In golang when we encode to json there is always a new line...
	expected := `{"error":"Route requested not found"}` + "\n"

	// when
	h := kithttp.MakeHealthCheckHandler(s, logger)
	h.ServeHTTP(r, req)

	// then
	assert.Equal(t, expected, r.Body.String())
	assert.Equal(t, http.StatusNotFound, r.Result().StatusCode)
}
