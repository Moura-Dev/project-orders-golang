PKGS ?= $(shell go list ./...)
POSTGRESQL_URL ?= postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
.PHONY: all services clean

unit-test:
	go test -covermode=atomic -coverprofile coverage.out -v ${PKGS}  -short -tags=unit

integration-test:
	go test -covermode=atomic -coverprofile coverage.out -v ${PKGS} -tags=integration 

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

seed:
	docker exec orderbook-psql psql -d orderbook -f /tmp/seed.sql

users:
	docker exec orderbook-psql psql -d orderbook -f /tmp/users.sql

down:
	docker-compose down && docker volume prune -f

ensure-dependencies:
	go mod tidy

services-local:
	make services
	docker exec orderbook-psql sh /tmp/health.sh
	make users
	make migrate_up
	make seed

services:
	docker-compose up -d psql zoo redis-server kafka
