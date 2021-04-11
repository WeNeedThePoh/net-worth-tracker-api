package health

import (
	"github.com/jmoiron/sqlx"
)

// Service describes the behavior of a health service
type Service interface {
	CheckHealth() string
}

// healthService is an implementation of the health service interface
type healthService struct {
	db *sqlx.DB
}

// NewService returns a new userService
func NewService(database *sqlx.DB) Service {
	return &healthService{
		db: database,
	}
}

func (s *healthService) CheckHealth() string {
	var (
		dbStatus string = "OK"
	)

	err := s.db.Ping()
	if err != nil {
		dbStatus = "Something wrong with database connection"
	}

	return dbStatus
}
