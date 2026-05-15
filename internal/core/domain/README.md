# Domain Layer

Pure business domain — **zero external dependencies**. This package must NOT import from `adapter/`, `app/`, `infrastructure/`, or any framework.

## Structure

```
domain/
└── entity/        → Domain entities, value objects, events, enums
```

All domain types live in `entity/` for simplicity. Separate into subpackages (`valueobject/`, `event/`, `enum/`) only when the package grows too large.

## entity/

Contains all domain types:

- **Entities** — Structs with identity (ID)
- **Value Objects** — Structs compared by value (Money, Address)
- **Domain Events** — Structs representing something that happened (OrderCreated)
- **Enums** — Typed constants with `iota` (OrderStatus, PaymentMethod)

```go
package entity

type User struct {
    ID    string
    Name  string
    Email string
}

type CheckHealth struct {
    MongoDB bool
    Redis   bool
}
```

## Rules

- **No external imports** — Only standard library allowed
- **No business logic** — Logic lives in `app/` services, not here
- If this package grows too large (20+ files), consider splitting into `valueobject/`, `event/`, `enum/`
