generate: di-generate mock-clean mock-generate fmt

fmt:
	@echo "Formatting code..."

	go fmt ./...

	@echo "Code formatting complete."

start:
	@echo "Starting the application..."

	go run ./cmd/api/main.go

test:
	@echo "Running tests..."
	go test -v ./... -cover
	@echo "Tests complete."

mock-clean:
	@echo "Cleaning up mock files..."
	find . -type d -name "mock" -exec rm -rf {} +
	@echo "Mock files cleaned."

mock-generate:
	@echo "Generating mock files..."
	mockery
	@echo "Mock files generation complete."

di-generate:
	@echo "Generating dependency injection code..."
	wiresetgen generate

	go fmt ./cmd/api/di

	wire ./...
	@echo "Dependency injection code generation complete."