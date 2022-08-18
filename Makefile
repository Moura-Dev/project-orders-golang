POSTGRESQL_URL ?= postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: all services clean
# migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path db/migrations up
migrate_up:
    migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
    migrate -database ${POSTGRESQL_URL} -path db/migrations down

populate:
    docker exec orderbook-psql psql -d orderbook -f /tmp/seed.sql

users:
    docker exec orderbook-psql psql -d orderbook -f /tmp/users.sql

migrate_create:
	migrate create -dir db/migrations  -ext sql -seq create_tables




unit-test:
    go test -covermode=atomic -coverprofile coverage.out -v ${PKGS}  -short -tags=unit

integration-test:
    go test -covermode=atomic -coverprofile coverage.out -v ${PKGS} -tags=integration 


users:
    docker exec orderbook-psql psql -d orderbook -f /tmp/users.sql

down:
    docker-compose down && docker volume prune -f

tidy:
    go mod tidy

run:
	go run main.go

services-local:
    make services
    docker exec orderbook-psql sh /tmp/health.sh
    make users
    make migrate_up
    make seed

services:
    docker-compose up -d psql zoo redis-server kafka