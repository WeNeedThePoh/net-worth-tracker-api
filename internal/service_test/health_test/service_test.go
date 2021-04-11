package health_test

import (
	"os"
	"testing"

	"weneedthepoh/net-worth-tracker-api/internal/service/health"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gotest.tools/assert"
)

func TestHealthCheck(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Mock DB pool connection
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	s := health.NewService(db)

	// when
	dbStatus := s.CheckHealth()

	// then
	assert.Equal(t, "OK", dbStatus)
}

func TestHealthCheckOnDbError(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Mock DB pool connection and force to close DB connection straight away so the ping returns error
	mockDB, _, _ := sqlmock.New()
	mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	s := health.NewService(db)

	// when
	dbStatus := s.CheckHealth()

	// then
	assert.Equal(t, "Something wrong with database connection", dbStatus)
}
