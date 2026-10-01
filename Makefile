DB_NAME=lms_db
DB_URL=postgres://postgres:postgres@localhost:5432/$(DB_NAME)?sslmode=disable

.PHONY: run sqlc migrate-up migrate-down db-create setup tidy

db-create:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | findstr 1 >nul || psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

run:
	go run ./cmd/api

sqlc:
	sqlc generate

migrate-up:
	goose -dir internal/adapters/postgresql/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir internal/adapters/postgresql/migrations postgres "$(DB_URL)" down

db-create:
	psql -U postgres -h localhost -p 5432 -c "CREATE DATABASE lms_db;"

setup:
	$(MAKE) db-create
	$(MAKE) migrate-up
	$(MAKE) sqlc
	$(MAKE) run

tidy:
	go mod tidy

dev:
	air