.PHONY: help run build clean deps fmt vet

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run the server
	go run cmd/server/main.go

build: ## Build binary
	go build -o flashlight-go cmd/server/main.go

clean: ## Clean build artifacts
	rm -f flashlight-go
	rm -f coverage.out coverage.html

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

db-setup: ## Create database
	createdb flashlight_db

db-drop: ## Drop database
	dropdb flashlight_db
