run-rest:
	@echo "Start running REST API"
	@go run cmd/main.go rest

vendor:
	@echo "Downloading dependencies..."
	@go mod vendor