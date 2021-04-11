# We need to this because in production we already have the env variables difined
ifndef DATABASE_URL
	DATABASE_URL = postgres://root:rootadmin@localhost:5433/net_worth?sslmode=disable
endif

MIGRATION_COMMAND = migrate

build:
	go build -o bin/server cmd/server/main.go

prepare_config:
	cp configs/.config.prod.yaml configs/.config.yaml
	sed -i -e "s/DB_HOST/${DB_HOST}/g" ./configs/.config.yaml
	sed -i -e "s/DB_PORT/${DB_PORT}/g" ./configs/.config.yaml
	sed -i -e "s/DB_NAME/${DB_NAME}/g" ./configs/.config.yaml
	sed -i -e "s/DB_USER/${DB_USER}/g" ./configs/.config.yaml
	sed -i -e "s/DB_PASSWORD/${DB_PASSWORD}/g" ./configs/.config.yaml
	sed -i -e "s/DB_SSL/${DB_SSL}/g" ./configs/.config.yaml
	sed -i -e "s/PORT/${PORT}/g" ./configs/.config.yaml

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
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
	chmod +x migrate.linux-amd64
	$(eval MIGRATION_COMMAND = ./migrate.linux-amd64)
endif

	${MIGRATION_COMMAND} -database ${DATABASE_URL} -path db/migrations up

migrate_down:
	migrate -database ${DATABASE_URL} -path db/migrations down
