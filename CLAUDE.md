# CLAUDE.md

Project conventions and guidelines for AI assistants working on this codebase.

## Project Overview

Go service boilerplate using **Hexagonal Architecture** (Ports and Adapters) with Cobra CLI, Echo HTTP framework, MongoDB, and Redis.

## Architecture Rules

### Layer Boundaries (STRICT)

1. **`internal/core/`** — Pure domain. MUST NOT import from `adapter/`, `app/`, `infrastructure/`, or any external framework.
2. **`internal/app/`** — Application services. Depends ONLY on `core/` ports and entities.
3. **`internal/adapter/`** — Implements ports. Can import `core/`, `pkg/`, and external libraries.
4. **`internal/infrastructure/`** — Technical setup (DB clients, logger). No business logic.
5. **`cmd/`** — Composition Root. Wires everything together. The ONLY place for dependency injection.

### Dependency Direction

```
cmd/ → adapter/ → app/ → core/
         ↓
    infrastructure/
```

Dependencies ALWAYS flow inward. Never reference outer layers from inner layers.

## Conventions

### Naming

- **Ports**: Interface names in `core/port/inbound/` and `core/port/outbound/`.
- **Adapters**: Concrete implementations in `adapter/{inbound,outbound}/{concern}/{technology}/{feature}/`.
- **Services**: Application use cases in `app/{domain}/`.
- **DTOs**: Request/Response types in `adapter/inbound/rest/dto/`.

### Error Handling

- Use `internal/pkg/errs` for domain errors (`DomainError` with `ErrorCode`).
- Available codes: `NOT_FOUND`, `CONFLICT`, `UNAUTHORIZED`, `FORBIDDEN`, `VALIDATION`, `INTERNAL`.
- Constructors: `errs.NewNotFound()`, `errs.NewConflict()`, `errs.NewInternal(err, msg)`, etc.
- HTTP translation happens in `adapter/inbound/rest/utils/errmap.go` via `MapErrorToHTTP()`.
- NEVER return HTTP status codes from `core/` or `app/` layers.
- Use `errs.GetCode(err)` to check error types (e.g., `errs.GetCode(err) == errs.ErrNotFound`) instead of `errors.Is()` against sentinel errors, since outbound adapters now return typed `DomainError`.

### Mock Generation

- Mocks use `go.uber.org/mock` (gomock).
- Generated via `//go:generate mockgen` directives on port interface files.
- Mocks live in `mocks/` subdirectory next to each port file.
- Regenerate: `go generate ./internal/core/port/...`
- In tests, import with aliases: `repoMocks "...outbound/repository/mocks"`, `cacheMocks "...outbound/cache/mocks"`

### Testing

- Use **table-driven tests** with `[]struct` and `for _, tt := range tests`.
- Use `testify/assert` for assertions.
- Use `assertFn func(t *testing.T, got *Type, err error)` pattern in test cases.
- Mock dependencies via `testDeps` struct with `setupTest(t)` helper.
- Package-level tests (same package, not `_test` suffix) for access to unexported types.

### Configuration

- Environment variables loaded from `.env` via `internal/pkg/config`.
- Access config values through typed functions: `config.AppPort()`, `config.MongoDBName()`, etc.
- For unit tests, mock env or use test-specific `.env` files.

## Build & Run

```bash
make install          # go mod tidy && go mod vendor
make docker-up        # start all services via Docker Compose
make docker-down      # stop docker services
make run-rest         # go run cmd/main.go rest
go test ./...         # run all tests
go build ./...        # verify compilation
go vet ./...          # static analysis
```

## Adding a New Feature

1. Define domain entity in `internal/core/domain/entity/`.
2. Define inbound port (service interface) in `internal/core/port/inbound/`.
3. Define outbound port (repository interface) in `internal/core/port/outbound/`.
4. Implement application service in `internal/app/{feature}/`.
5. Implement outbound adapter in `internal/adapter/outbound/{concern}/{tech}/{feature}/`.
6. Implement inbound adapter (handler) in `internal/adapter/inbound/rest/handler/{feature}/`.
7. Add DTO in `internal/adapter/inbound/rest/dto/`.
8. Register route in `internal/adapter/inbound/rest/router/router.go`.
9. Wire dependencies in `cmd/api/rest/cmd.go` (Composition Root).
10. Generate mocks: `go generate ./internal/core/port/...`
11. Write tests in `internal/app/{feature}/`.

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/labstack/echo/v4` | HTTP framework |
| `github.com/spf13/cobra` | CLI framework |
| `github.com/spf13/viper` | Configuration |
| `go.mongodb.org/mongo-driver` | MongoDB driver |
| `github.com/redis/go-redis/v9` | Redis client |
| `go.uber.org/mock` | Mock generation |
| `github.com/stretchr/testify` | Test assertions |
| `github.com/go-ozzo/ozzo-validation/v4` | DTO validation |
| `github.com/golang-jwt/jwt/v5` | JWT generation & validation |
| `golang.org/x/sync/errgroup` | Concurrent error handling |
| `golang.org/x/crypto/bcrypt` | Password hashing |
