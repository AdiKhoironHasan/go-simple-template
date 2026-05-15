# Repository Adapters

Outbound adapters for **persistent data storage** (databases).

## Structure

```
repository/
├── mongo/              # MongoDB implementations
│   └── <feature>/      # e.g., health/, user/, order/
│       ├── repository.go       # Constructor + struct
│       └── <method>.go         # One file per method
├── postgres/           # PostgreSQL implementations
│   └── <feature>/
│       ├── repository.go
│       └── <method>.go
└── README.md
```

## Implementation Guide

Each feature sub-package implements one `outbound.XxxRepository` port interface.

```go
package health

import (
    "go-simple-template/internal/core/port/outbound"
    "go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
    client *mongo.Client
}

func New(client *mongo.Client) outbound.HealthRepository {
    return &repository{client: client}
}
```

## Rules

- One feature = one sub-package under the technology folder
- Constructor returns the **port interface**, not a concrete type
- Package name reflects the **feature**, not the technology
- No cross-technology imports within this folder
