# Kafka Producer (Outbound Adapter)

## Purpose

Driven adapter that implements outbound port interfaces by **producing messages** to Kafka topics. Used for event streaming, inter-service communication, and asynchronous data pipelines.

## When to Use

- High-throughput event streaming (logs, metrics, audit trails).
- Event sourcing — persisting domain events to an append-only log.
- Data pipeline ingestion.
- Cross-service async communication with ordering guarantees.

## Folder Structure

```
kafka/
└── {feature}/
    ├── {feature}.go        → Constructor & struct (implements outbound port)
    └── produce.go          → Producer operations
```

## Implementation Guide

### 1. Define Outbound Port (in `core/port/outbound/`)

```go
// internal/core/port/outbound/audit_logger.go
package outbound

import "context"

type AuditLogger interface {
    LogAction(ctx context.Context, userID, action, resource string) error
}
```

### 2. Implement Adapter

```go
// kafka/audit/audit.go
package audit

import (
    "go-simple-template/internal/core/port/outbound"
    "github.com/segmentio/kafka-go"
)

type producer struct {
    writer *kafka.Writer
}

func New(writer *kafka.Writer) outbound.AuditLogger {
    return &producer{writer: writer}
}
```

```go
// kafka/audit/produce.go
package audit

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
)

func (p *producer) LogAction(ctx context.Context, userID, action, resource string) error {
    payload, _ := json.Marshal(map[string]string{
        "user_id":  userID,
        "action":   action,
        "resource": resource,
    })
    return p.writer.WriteMessages(ctx, kafka.Message{
        Key:   []byte(userID),
        Value: payload,
    })
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
kafkaWriter := infrastructure.NewKafkaWriter(cfg, "audit-log-topic")
auditLogger := kafkaAudit.New(kafkaWriter)
userService := userApp.New(userRepo, auditLogger)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/segmentio/kafka-go` | Kafka client (pure Go) |
| `github.com/confluentinc/confluent-kafka-go` | Alternative: librdkafka-based client |

## Architecture Rules

- Implementations MUST satisfy `core/port/outbound` interfaces.
- Use partition keys for ordering guarantees where needed.
- The domain layer produces via the port interface — no Kafka imports in `core/`.
