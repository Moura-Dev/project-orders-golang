POSTGRESQL_URL ?= postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: all services clean
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