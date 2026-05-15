# Kafka Consumer (Inbound Adapter)

## Purpose

Driving adapter that consumes messages from **Apache Kafka** topics. Translates incoming events/messages into calls to inbound port interfaces, enabling event-driven and streaming architectures.

## When to Use

- High-throughput event streaming (logs, metrics, clickstreams).
- Event sourcing and CQRS read-model projections.
- Cross-service data synchronization.
- Ordered message processing with partition-based parallelism.

## Folder Structure

```
kafka/
├── consumer.go         → Kafka consumer group setup (broker connection, topic subscription)
├── handler/            → Message handlers (maps topic messages to port/inbound calls)
│   └── {feature}/      → Feature-specific topic handler
└── dto/                → Message payload structures
```

## Implementation Guide

### 1. Implement Consumer Handler

```go
// handler/analytics/handler.go
package analytics

import (
    "context"
    "encoding/json"

    "go-simple-template/internal/core/port/inbound"

    "github.com/segmentio/kafka-go"
)

type Handler struct {
    service inbound.AnalyticsService
}

func New(service inbound.AnalyticsService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, msg kafka.Message) error {
    var event ClickEvent
    if err := json.Unmarshal(msg.Value, &event); err != nil {
        return err
    }
    return h.service.TrackEvent(ctx, event.UserID, event.Action)
}
```

### 2. Wire in Composition Root (`cmd/`)

```go
kafkaReader := infrastructure.NewKafkaReader(cfg, "analytics-topic", "consumer-group-1")
analyticsHandler := kafkaAnalyticsHandler.New(analyticsService)
consumer := kafkaConsumer.New(kafkaReader, analyticsHandler)
go consumer.Start(ctx)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/segmentio/kafka-go` | Kafka client (pure Go) |
| `github.com/confluentinc/confluent-kafka-go` | Alternative: librdkafka-based client |

## Architecture Rules

- Handlers MUST depend on `core/port/inbound` interfaces.
- Consumer group management is an infrastructure concern — keep it in `consumer.go`.
- Offset commit strategy should be configurable.
