# CLAUDE.md

Project conventions and guidelines for AI assistants working on this codebase.

## Project Overview

Go service boilerplate using **Hexagonal Architecture** (Ports and Adapters) with Cobra CLI, Echo HTTP framework, MongoDB, and Redis.

## Architecture Rules

### Layer Boundaries (STRICT)

1. **`internal/core/`** — Pure domain (entities, ports, domain services). MUST NOT import from any other layer or external framework.
2. **`internal/service/`** — Application services (use cases). Depends on `core/` ports and entities, plus `pkg/` utilities. MUST NOT import from `adapter/` or `infrastructure/`.
3. **`internal/adapter/`** — Implements ports (inbound/outbound). Can import `core/`, `service/`, `pkg/`, and external libraries.
4. **`internal/infrastructure/`** — Technical setup (DB clients, logger). No business logic.
5. **`cmd/`** — Composition Root. Wires everything together. The ONLY place for dependency injection.

### Dependency Direction

```
cmd/ ──→ adapter/ ──→ service/ ──→ core/
  │         │            │
  └─────────┴────────────┴──→ infrastructure/  pkg/
```

Dependencies ALWAYS flow inward. Never reference outer layers from inner layers.

- `cmd/` imports `adapter/`, `infrastructure/`, `service/`, and `pkg/` (composition root).
- `adapter/` imports `service/` (inbound handlers) or implements ports directly. Can import `pkg/`.
- `service/` imports `core/port/`, `core/domain/entity/`, and `pkg/`.
- `core/` imports nothing outside itself.

### pkg/ Usage Rules

`internal/pkg/` is for shared internal utilities — pure functions, no I/O, no side effects. These are stateless helpers that happen to live outside the domain but don't require port/interface abstraction.

- **`pkg/errs`**, `pkg/consts` — allowed everywhere.
- **`pkg/jwt`** (token generation, validation) — pure in-memory crypto operations. Allowed in `service/` and `adapter/`. No I/O, no external dependency beyond `golang-jwt/jwt`.
- **`pkg/crypto`** (password hashing, bcrypt) — pure in-memory crypto operations. Allowed in `service/` and `adapter/`.
- **`pkg/config`** — environment variable loader. Allowed everywhere.
- **`pkg/context`** — request context helpers. Allowed everywhere.

**Rule of thumb:** If a pkg utility involves I/O (network call, filesystem, external system), it belongs in `adapter/outbound/` via a port. If it's purely in-memory computation, it stays in `pkg/`.

> **Design decision:** JWT and bcrypt are kept in `pkg/` — not extracted to outbound ports. Rationale: these are stateless, in-memory utilities that don't involve external I/O. Abstracting them behind ports adds indirection without solving a real problem. See `.claude/my.md` for the full argument.

### Linter Enforcement (Recommended)

Use `depguard` in `.golangci.yml` to enforce layer boundaries:

```yaml
version: "2"
linters:
  settings:
    depguard:
      rules:
        service-layer:
          files:
            - $all
            - "!**/cmd/**"
            - "!**/adapter/**"
            - "!**/infrastructure/**"
            - "!**/pkg/**"
            - "!**/mocks/**"
          deny:
            - pkg: "go-simple-template/internal/adapter"
              desc: "service and core layers must not import adapter"
            - pkg: "go-simple-template/internal/infrastructure"
              desc: "service and core layers must not import infrastructure"
```

## Conventions

### Naming

- **Ports**: Interface names in `core/port/inbound/` and `core/port/outbound/`.
- **Inbound ports**: `{Domain}Service` (e.g., `AuthService`, `HealthService`).
- **Outbound ports**: `{Concern}` (e.g., `UserRepository`, `Health`).
- **Adapters**: Concrete implementations in `adapter/{inbound,outbound}/{concern}/{technology}/{feature}/`.
- **Application Services**: Implement inbound ports, live in `service/{domain}/`.
- **DTOs**: Request/Response types in `adapter/inbound/rest/dto/`.

### Error Handling

- Use `internal/pkg/errs` for domain errors (`DomainError` with `ErrorCode`).
- Available codes: `NOT_FOUND`, `CONFLICT`, `UNAUTHORIZED`, `FORBIDDEN`, `VALIDATION`, `INTERNAL`.
- Constructors: `errs.NewNotFound()`, `errs.NewConflict()`, `errs.NewInternal(err, msg)`, etc.
- HTTP translation happens in `adapter/inbound/rest/utils/errmap.go` via `MapErrorToHTTP()`.
- NEVER return HTTP status codes from `core/` or `service/` layers.
- Use `errs.GetCode(err)` to check error types (e.g., `errs.GetCode(err) == errs.ErrNotFound`) instead of `errors.Is()` against sentinel errors.

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
make verify           # generate mocks + build + vet + lint + test (all-in-one)
make generate-mocks   # go generate ./internal/core/port/...
make build            # go build ./...
make vet              # go vet ./...
make lint             # golangci-lint run ./...
make test             # go test ./...
```

## Adding a New Feature

1. Define domain entity in `internal/core/domain/entity/`.
2. Define inbound port (service interface) in `internal/core/port/inbound/`.
3. Define outbound ports (repository, etc.) in `internal/core/port/outbound/` — only for I/O dependencies.
4. Implement service in `internal/service/{feature}/` — depend on ports, entities, and pkg/ utilities.
5. Implement outbound adapter in `internal/adapter/outbound/{concern}/{tech}/{feature}/`.
6. Implement inbound adapter (handler) in `internal/adapter/inbound/rest/handler/{feature}/`.
7. Add DTO in `internal/adapter/inbound/rest/dto/`.
8. Register route in `internal/adapter/inbound/rest/router/router.go`.
9. Wire dependencies in `cmd/api/rest/cmd.go` (Composition Root).
10. Generate mocks: `go generate ./internal/core/port/...`
11. Write tests in `internal/service/{feature}/`.

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
