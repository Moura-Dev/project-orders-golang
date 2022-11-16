include .env
.PHONY: all services clean

migrate_up:
	docker run -v $(shell pwd)/db/migrations:/migrations --network host migrate/migrate -path migrations -database ${DB_URL} up

migrate_down:
	docker run -v $(shell pwd)/db/migrations:/migrations --network host migrate/migrate -path migrations -database ${DB_URL} down -all

down:
	docker-compose down && docker volume prune -f

up:
	docker-compose up -d
