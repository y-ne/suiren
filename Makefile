include .env

BINARY_NAME=suiren
MAIN_FILE=cmd/api/main.go

.PHONY: migrate-up migrate-down migrate-create migrate-force migrate-fix install-tools sqlc run build

build:
	go build -o bin/${BINARY_NAME} ${MAIN_FILE}

run:
	go run ${MAIN_FILE}

sqlc:
	sqlc generate

install-tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create new migration
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migrations -seq $$name

# Run migrations up
migrate-up:
	migrate -path internal/db/migrations -database "$(DATABASE_URL)" up

# Run migrations down
migrate-down:
	migrate -path internal/db/migrations -database "$(DATABASE_URL)" down

# Force specific version
migrate-force:
	@read -p "Enter version: " version; \
	migrate -path internal/db/migrations -database "$(DATABASE_URL)" force $$version

# Fix dirty state
migrate-fix:
	migrate -path internal/db/migrations -database "$(DATABASE_URL)" force $(shell migrate -path internal/db/migrations -database "$(DATABASE_URL)" version)