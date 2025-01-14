include .env
export

migrate-up:
	migrate -database "$(DATABASE_URL)" -path migrations up

migrate-down:
	migrate -database "$(DATABASE_URL)" -path migrations down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

run:
	go run cmd/api/main.go