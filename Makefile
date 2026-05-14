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