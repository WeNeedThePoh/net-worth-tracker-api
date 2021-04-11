package main

import (
	"fmt"
	"os"

	"weneedthepoh/net-worth-tracker-api/internal/config"
	httpServer "weneedthepoh/net-worth-tracker-api/internal/transport/http"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	err := config.InitFromFile("configs/.config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	configs := config.GetAll()
	db := openDB(configs)

	httpServer.StartServer(configs, db, logger)
}

func openDB(configs config.Conf) *sqlx.DB {
	var db *sqlx.DB
	var err error

	dataSource := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", configs.Db.User, configs.Db.Password, configs.Db.Host, configs.Db.Port, configs.Db.Name, configs.Db.Ssl)
	db, err = sqlx.Connect("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	return db
}
