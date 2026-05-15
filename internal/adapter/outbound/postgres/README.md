# PostgreSQL Outbound Adapter

## Purpose

Driven adapter that implements outbound port interfaces using **PostgreSQL** as the persistence layer. Handles data access, queries, and transactions for domain entities.

## When to Use

- Relational data models with complex relationships and joins.
- ACID transaction requirements.
- Data requiring strong consistency guarantees.
- Reporting and analytics queries.

## Folder Structure

```
postgres/
└── {feature}/
    ├── {feature}.go        → Constructor & struct definition (implements outbound port)
    ├── create.go           → Insert operations
    ├── find.go             → Query/read operations
    ├── update.go           → Update operations
    └── delete.go           → Delete operations
```

> Follow the same file-per-operation pattern as the existing `mongo/` adapter.

## Implementation Guide

### 1. Define Outbound Port (in `core/port/outbound/`)

```go
// internal/core/port/outbound/user_repository.go
package outbound

import (
    "context"
    "go-simple-template/internal/core/domain/entity"
)

type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    FindByID(ctx context.Context, id string) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id string) error
}
```

### 2. Implement Adapter

```go
// postgres/user/user.go
package user

import (
    "go-simple-template/internal/core/port/outbound"

    "github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
    pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) outbound.UserRepository {
    return &repository{pool: pool}
}
```

```go
// postgres/user/find.go
package user

import (
    "context"
    "go-simple-template/internal/core/domain/entity"
    "go-simple-template/internal/pkg/errs"
)

func (r *repository) FindByID(ctx context.Context, id string) (*entity.User, error) {
    var user entity.User
    err := r.pool.QueryRow(ctx,
        "SELECT id, name, email FROM users WHERE id = $1", id,
    ).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, errs.NewNotFound("user", id)
    }
    return &user, nil
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
pgPool := infrastructure.NewPostgres(cfg)
userRepo := pgUser.New(pgPool)
userService := userApp.New(userRepo)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/jackc/pgx/v5` | PostgreSQL driver (recommended) |
| `github.com/jackc/pgx/v5/pgxpool` | Connection pooling |
| `github.com/lib/pq` | Alternative: standard `database/sql` driver |
| `github.com/golang-migrate/migrate` | Database migrations |

## Infrastructure Setup

The corresponding PostgreSQL client setup belongs in `internal/infrastructure/postgres.go`:

```go
func NewPostgres(cfg config.PostgresConfig) *pgxpool.Pool {
    pool, err := pgxpool.New(context.Background(), cfg.DSN())
    if err != nil {
        log.Fatalf("failed to connect to postgres: %v", err)
    }
    return pool
}
```

## Architecture Rules

- Implementations MUST satisfy `core/port/outbound` interfaces.
- Return domain entities (`core/domain/entity`), NOT database-specific models.
- Map domain errors (`pkg/errs`) — never expose `pgx` errors to upper layers.
- SQL queries live in the adapter, not in `core/` or `app/`.
