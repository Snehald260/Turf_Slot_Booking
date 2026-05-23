.PHONY: build run test clean migrate-up migrate-down docker-up docker-down swagger

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test ./... -v -cover

# Run tests with race detector
test-race:
	go test ./... -v -race

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Database migrations
migrate-up:
	migrate -path migrations -database "postgresql://turf_user:turf_password@localhost:5432/turf_booking?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://turf_user:turf_password@localhost:5432/turf_booking?sslmode=disable" down

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Generate Swagger docs
swagger:
	swag init -g cmd/server/main.go -o docs/

# Seed database
seed:
	psql -h localhost -U turf_user -d turf_booking -f scripts/seed.sql

# Lint
lint:
	golangci-lint run ./...
