package endpoint_test

import (
	"context"
	"os"
	"testing"

	"weneedthepoh/net-worth-tracker-api/internal/endpoint"
	"weneedthepoh/net-worth-tracker-api/internal/service/health"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gotest.tools/assert"
)

func TestMakeHealthCheckEndpoint(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Mock DB pool connection
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	s := health.NewService(db)
	e := endpoint.MakeHealthCheckEndpoint(s)

	ctx := context.TODO()

	expected := endpoint.HealthCheckResponse{
		Database: "OK",
	}

	// when
	res, err := e(ctx, endpoint.HealthCheckRequest{})

	// then
	require.Nil(t, err)
	assert.Equal(t, expected, res)
}
