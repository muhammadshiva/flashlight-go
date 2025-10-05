# Flashlight GO Architecture

A production-ready Go project skeleton implementing clean architecture principles with modern tech stack.

## Architecture Overview

This project follows clean architecture with the following flow:

1. **External System** → Request (HTTP, gRPC, Messaging)
2. **Delivery Layer** → Creates Models from request data
3. **Delivery Layer** → Calls Use Case with Model data
4. **Use Case** → Creates Entity for business logic
5. **Use Case** → Calls Repository with Entity data
6. **Repository** → Performs database operations
7. **Use Case** → Creates Models for Gateway or from Entity
8. **Use Case** → Calls Gateway with Model data
9. **Gateway** → Constructs request to external system
10. **Gateway** → Performs external API calls

## Tech Stack

- **Language**: Go 1.21+
- **Database**: MySQL
- **Message Queue**: Apache Kafka

## Framework & Libraries

- **GoFiber** - HTTP Framework
- **GORM** - ORM
- **Viper** - Configuration Management
- **Golang Migrate** - Database Migration
- **Go Playground Validator** - Request Validation
- **Logrus** - Structured Logging
- **Sarama** - Kafka Client

## Project Structure

```
.
├── api/                    # API specifications (OpenAPI/Swagger)
├── cmd/
│   ├── web/               # Web server entry point
│   └── worker/            # Worker/consumer entry point
├── db/
│   └── migrations/        # Database migration files
├── internal/
│   ├── delivery/          # Delivery layer (HTTP, gRPC, Kafka)
│   │   ├── http/
│   │   ├── grpc/
│   │   └── kafka/
│   ├── entity/            # Domain entities
│   ├── gateway/           # External service integrations
│   ├── model/             # Request/Response models
│   ├── repository/        # Data access layer
│   └── usecase/           # Business logic
├── pkg/                   # Shared packages
│   ├── config/
│   ├── database/
│   ├── kafka/
│   ├── logger/
│   └── validator/
├── test/                  # Test files
├── config.json            # Application configuration
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- MySQL 8.0+
- Apache Kafka (optional, for worker)
- golang-migrate CLI

Install golang-migrate:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Windows (using scoop)
scoop install migrate
```

### Configuration

Edit `config.json` to match your environment:

```json
{
  "app": {
    "name": "Go Clean Architecture",
    "env": "development",
    "port": 3000
  },
  "database": {
    "host": "localhost",
    "port": 3306,
    "username": "root",
    "password": "",
    "name": "golang_clean_architecture"
  },
  "kafka": {
    "brokers": ["localhost:9092"],
    "consumer_group": "go-skeleton-consumer"
  }
}
```

### Database Setup

1. Create database:
```bash
mysql -u root -p -e "CREATE DATABASE golang_clean_architecture;"
```

2. Create migration:
```bash
migrate create -ext sql -dir db/migrations create_table_examples
```

3. Run migrations:
```bash
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Run tests:
```bash
go test -v ./test/
```

### Running the Application

#### Web Server (HTTP/gRPC)

```bash
go run cmd/web/main.go
```

The server will start on `http://localhost:3000`

Health check: `GET http://localhost:3000/health`

#### Worker (Kafka Consumer)

```bash
go run cmd/worker/main.go
```

## API Endpoints

### Examples API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/examples` | Create new example |
| GET | `/api/v1/examples` | Get all examples (with pagination) |
| GET | `/api/v1/examples/:id` | Get example by ID |
| PUT | `/api/v1/examples/:id` | Update example |
| DELETE | `/api/v1/examples/:id` | Delete example |

See `api/openapi.yaml` for complete API specification.

## Development

### Adding New Features

1. **Create Entity** in `internal/entity/`
2. **Create Model** in `internal/model/`
3. **Create Repository** interface and implementation in `internal/repository/`
4. **Create Gateway** (if external API is needed) in `internal/gateway/`
5. **Create Use Case** in `internal/usecase/`
6. **Create Delivery Handler** in `internal/delivery/http/` or `internal/delivery/kafka/`
7. **Update Router** in `internal/delivery/http/router.go`
8. **Write Tests** in `test/`

### Migration Commands

Create migration:
```bash
migrate create -ext sql -dir db/migrations <migration_name>
```

Run migrations:
```bash
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

Rollback migration:
```bash
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations down
```

## Testing

Run all tests:
```bash
go test -v ./test/
```

Run with coverage:
```bash
go test -v -cover ./test/
```

## Building for Production

```bash
# Build web server
go build -o bin/web cmd/web/main.go

# Build worker
go build -o bin/worker cmd/worker/main.go
```

## Contributing

1. Follow the clean architecture principles
2. Write tests for new features
3. Update API documentation
4. Follow Go best practices and conventions

## License

MIT License
