.PHONY: help test run-web run-worker build clean migrate-up migrate-down migrate-create

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test -v ./test/

test-cover: ## Run tests with coverage
	go test -v -cover -coverprofile=coverage.out ./test/
	go tool cover -html=coverage.out -o coverage.html

run-web: ## Run web server
	go run cmd/web/main.go

run-worker: ## Run worker
	go run cmd/worker/main.go

build: ## Build binaries
	go build -o bin/web cmd/web/main.go
	go build -o bin/worker cmd/worker/main.go

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

migrate-up: ## Run database migrations
	migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up

migrate-down: ## Rollback last migration
	migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations down 1

migrate-create: ## Create new migration (usage: make migrate-create name=create_users_table)
	migrate create -ext sql -dir db/migrations $(name)

deps: ## Download dependencies
	go mod download
	go mod tidy

lint: ## Run linter (requires golangci-lint)
	golangci-lint run

fmt: ## Format code
	go fmt ./...
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...
