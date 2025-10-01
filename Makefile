.PHONY: help run build test clean install migrate

help: ## Show this help message
	@echo "Available commands:"
	@echo "  make run       - Run the application in development mode"
	@echo "  make build     - Build the application binary"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make install   - Install dependencies"

run: ## Run the application
	go run cmd/api/main.go

build: ## Build the application
	go build -o bin/api.exe cmd/api/main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build artifacts
	rm -rf bin/
	go clean

install: ## Install dependencies
	go mod download
	go mod tidy

dev: ## Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
	air
