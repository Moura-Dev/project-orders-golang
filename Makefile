include .env
.PHONY: all services clean

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

down:
	docker-compose down && docker volume prune -f

up:
	docker-compose up -d
