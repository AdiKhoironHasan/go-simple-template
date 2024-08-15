start-rest:
	@echo "Starting REST API"
	@go run main.go rest

consumer-ping:
	@echo "Pinging consumer"
	@go run main.go consumer ping
