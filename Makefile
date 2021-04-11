# We need to this because in production we already have the env variables difined
ifndef DATABASE_URL
	DATABASE_URL = postgres://root:rootadmin@localhost:5433/net_worth?sslmode=disable
endif

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

build_docker:
	docker-compose up --build -d

test:
	go test ./...

test_unit:
	go test --short ./...

lint:
	go fmt ./...

create_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate:
ifeq (, $(shell which migrate))
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
endif

	migrate -database ${DATABASE_URL} -path db/migrations up

migrate_down:
	migrate -database ${DATABASE_URL} -path db/migrations down
