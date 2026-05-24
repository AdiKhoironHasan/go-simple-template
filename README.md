# go-simple-template

A production-ready Go service boilerplate following **Hexagonal Architecture** (Ports and Adapters). Designed for clean separation of concerns, testability, and scalability.

## Architecture

```
cmd/                          → Entrypoints & Composition Root
├── main.go                   → Cobra CLI bootstrap (registers all subcommands)
├── api/
│   ├── rest/cmd.go           → REST server wiring (DI)
│   └── grpc/cmd.go           → gRPC server wiring (DI)
├── consumer/
│   ├── rmq/cmd.go            → RabbitMQ consumer wiring
│   └── kafka/cmd.go          → Kafka consumer wiring
└── scheduler/cmd.go          → Cron scheduler wiring

internal/
├── core/                     → Domain (pure business logic, zero dependencies)
│   ├── domain/entity/        → Domain types (entities, value objects, events, enums)
│   └── port/                 → Interfaces (contracts)
│       ├── inbound/          → Driving ports (service interfaces)
│       └── outbound/         → Driven ports (concern-based)
│           ├── repository/   → Database port interfaces
│           └── cache/        → Cache port interfaces
│
├── service/                      → Application services (use cases)
│   ├── health/               → Health check use case
│   └── auth/                 → User authentication (register, login, refresh, logout, profile)
│
├── adapter/                  → Infrastructure adapters
│   ├── inbound/              → Driving adapters (entry points)
│   │   ├── rest/             → REST API (Echo)
│   │   │   ├── handler/      → HTTP handlers
│   │   │   ├── dto/          → Request/Response DTOs
│   │   │   ├── router/       → Route registration
│   │   │   ├── server/       → Echo server setup
│   │   │   ├── middleware/   → HTTP middlewares
│   │   │   └── utils/        → Error mapping (domain → HTTP)
│   │   ├── grpc/             → gRPC server & handlers
│   │   ├── consumer/         → Message consumers
│   │   │   ├── rmq/          → RabbitMQ consumer
│   │   │   └── kafka/        → Kafka consumer
│   │   └── scheduler/        → Cron/scheduled jobs
│   │
│   └── outbound/             → Driven adapters (concern-based)
│       ├── repository/       → Database adapters
│       │   ├── mongo/        → MongoDB (e.g., mongo/health/)
│       │   └── postgres/     → PostgreSQL
│       ├── cache/            → In-memory / cache decorators
│       │   ├── redis/        → Redis (e.g., redis/health/)
│       │   └── memcached/    → Memcached
│       ├── client/           → External service integrations
│       │   ├── rest/         → REST API clients
│       │   └── grpc/         → gRPC clients
│       ├── messaging/        → Event/message publishers
│       │   ├── kafka/        → Kafka producers
│       │   └── rabbitmq/     → RabbitMQ publishers
│       └── storage/          → File/object storage
│           ├── s3/           → AWS S3
│           └── gcs/          → Google Cloud Storage
│
├── infrastructure/           → Technical setup (DB clients, logger)
│   ├── mongodb.go
│   ├── redis.go
│   └── slog.go
│
└── pkg/                      → Shared internal packages
    ├── config/               → Environment configuration
    ├── consts/               → Constants
    ├── context/              → Context helpers (user context)
    ├── crypto/               → Crypto utilities (bcrypt, AES-GCM, SHA, RSA)
    ├── errs/                 → Domain error system
    └── jwt/                  → JWT generation & validation
```

### Flow

```
HTTP Request --> DTO --> Adapter(inbound) --> Port(inbound) --> Service(app) --> Port(outbound) --> Adapter(outbound)
```

All dependencies flow **inward**. The core layer has zero knowledge of transport or infrastructure.

## Prerequisites

- **Go** 1.26+
- **MongoDB**
- **Redis** 

## Getting Started

### 1. Clone & Configure

```bash
git clone <repo-url>
cd go-simple-template
cp .env.example .env
# Edit .env with your database credentials (or use Docker Compose defaults)
```

### 2. Run with Docker Compose (Recommended)

```bash
make docker-up
```
This starts the REST API, MongoDB, and Redis containers with the default environment.

### 3. Install Dependencies (For Local Dev)

```bash
make install
```

### 4. Run the Server (Local)

```bash
make run-rest
```

The local server starts on the port defined in `APP_PORT` (default: `8080`).

### 5. Verify

```bash
# Basic health
curl localhost:8080/healthz

# Full health check (MongoDB + Redis)
curl "localhost:8080/healthz?mongodb=true&redis=true"
```

## Environment Variables

| Variable | Description | Default |
|---|---|---|
| `APP_NAME` | Application name | `go-simple-template` |
| `APP_VERSION` | Application version | `1.0.0` |
| `APP_PORT` | HTTP server port | `8080` |
| `APP_DEBUG` | Enable debug logging | `true` |
| `APP_SECRET_KEY` | JWT access token signing key (HMAC) | *(required)* |
| `APP_REFRESH_KEY` | JWT refresh token signing key (HMAC) | *(required)* |
| `MONGODB_PROTOCOL` | MongoDB protocol | `mongodb` |
| `MONGODB_ADDRESS` | MongoDB host:port | `localhost:27017` |
| `MONGODB_USERNAME` | MongoDB username | |
| `MONGODB_PASSWORD` | MongoDB password | |
| `MONGODB_NAME` | Database name | `go-simple` |
| `MONGODB_OPTION` | Connection string options | |
| `MONGODB_MAX_CONN_OPEN` | Max open connections | `100` |
| `MONGODB_MAX_CONN_IDLE` | Max idle connections | `10` |
| `MONGODB_MAX_CONN_LIFETIME` | Connection max lifetime | `1h` |
| `REDIS_HOST` | Redis host | `localhost` |
| `REDIS_PORT` | Redis port | `6379` |
| `REDIS_USERNAME` | Redis username | |
| `REDIS_PASSWORD` | Redis password | |
| `REDIS_DB` | Redis database number | `0` |
| `REDIS_TLS` | Enable TLS | `false` |

## Makefile Commands

| Command | Description |
|---|---|
| `make install` | Download dependencies & vendor |
| `make run-rest` | Run REST API server locally |
| `make docker-up` | Start all services via Docker Compose |
| `make docker-down` | Stop all Docker services |

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package
go test -v ./internal/service/health/...
```

## Key Design Decisions

- **Hexagonal Architecture** — Domain logic is fully decoupled from transport and infrastructure.
- **Composition Root** — All dependency wiring happens explicitly in `cmd/api/rest/cmd.go`.
- **Domain Errors** — `internal/pkg/errs` provides protocol-agnostic error codes (`NOT_FOUND`, `CONFLICT`, etc.) translated to HTTP at the adapter layer.
- **Concurrent Health Checks** — Uses `errgroup` for parallel MongoDB + Redis ping.
- **Mock Generation** — Uses `go.uber.org/mock` with `//go:generate mockgen` directives on port interfaces.
- **JWT Rotation** — Refresh token rotation with old token blacklisting on each refresh.
- **Secure Token Storage** — SHA-256 hashed tokens in Redis keys (never stores raw JWTs).
- **Algorithm Enforcement** — JWT validation enforces HS256 only, rejects other signing methods.
- **Structured Logging** — JSON slog handler with automatic request ID injection.

## License

MIT
