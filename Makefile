.PHONY: all run-rest install docker-up docker-down verify generate-mocks build vet lint test

run-rest:
	@echo "Start running REST API"
	@go run cmd/main.go rest

install:
	@echo "Downloading dependencies..."
	@go mod tidy
	@go mod vendor

docker-up:
	@echo "Starting docker services..."
	@docker compose -f deployment/docker-compose.yml up --build -d

docker-down:
	@echo "Stopping docker services..."
	@docker compose -f deployment/docker-compose.yml down

## All-in-one verification: generate mocks, build, vet, and test.
verify: generate-mocks build vet lint test
	@echo "All checks passed."

generate-mocks:
	@echo "Generating mocks..."
	@go generate ./internal/core/port/...

build:
	@echo "Building..."
	@go build ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...

test:
	@echo "Running tests..."
	@go test ./...

