# We need to this because in production we already have the env variables difined
ifndef DB_HOST
	DB_HOST = localhost
endif

ifndef DB_PORT
	DB_PORT = 5433
endif

ifndef DB_NAME
	DB_NAME = net_worth
endif

ifndef DB_USER
	DB_USER = root
endif

ifndef DB_PASSWORD
	DB_PASSWORD = rootadmin
endif

DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

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

run_migration:
	migrate -database ${DB_URL} -path db/migrations up

down_migration:
	migrate -database ${DB_URL} -path db/migrations down
