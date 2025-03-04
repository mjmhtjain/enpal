.PHONY: build run test clean docker-build docker-run docker-stop help

# Default target
.DEFAULT_GOAL := help

# Go related variables
BINARY_NAME=enpal-app
MAIN_FILE=main.go

help: ## Display available commands
	@echo "Available commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run: build ## Build and run the application
	./$(BINARY_NAME)

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean: ## Clean build artifacts
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out

fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run linter
	if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed"; \
		exit 1; \
	fi

docker-build: ## Build Docker image
	docker-compose build

docker-run: ## Run application in Docker
	docker-compose up -d

docker-stop: ## Stop Docker containers
	docker-compose down

docker-rebuild: ## Rebuild Docker image
	docker-compose down
	docker-compose build

docker-cleanup: ## Clean up Docker containers and images
	docker-compose down
	docker image prune -f
	docker volume prune -f

docker-prune: ## Clean up Docker containers
	docker container prune -f

dev: ## Run application in development mode
	go run $(MAIN_FILE) 