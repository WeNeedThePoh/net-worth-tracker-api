package http

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"weneedthepoh/net-worth-tracker-api/internal/config"
	"weneedthepoh/net-worth-tracker-api/internal/service/health"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// StartServer Start new http server
func StartServer(configs config.Conf, db *sqlx.DB, logger log.Logger) {
	server := InitServer(configs, db, logger)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", server.Addr, "msg", "listening")
		errs <- http.ListenAndServe(server.Addr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

// InitServer Initialize the http server
func InitServer(configs config.Conf, db *sqlx.DB, logger log.Logger) http.Server {
	var (
		healthService health.Service = health.NewService(db)
	)

	r := mux.NewRouter()

	r.Handle("/health", MakeHealthCheckHandler(healthService, logger))
	http.Handle("/", handleCORS(r))

	r.NotFoundHandler = handleRouteNotFound(r)

	httpAddr := fmt.Sprintf(":%d", configs.Serve.Public.Port)

	return http.Server{
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
		Addr:         httpAddr,
	}
}
