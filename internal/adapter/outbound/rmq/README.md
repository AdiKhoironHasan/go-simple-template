# RabbitMQ Publisher (Outbound Adapter)

## Purpose

Driven adapter that implements outbound port interfaces by **publishing messages** to RabbitMQ exchanges. Used when the application needs to emit events or dispatch async tasks to other services or workers.

## When to Use

- Publishing domain events after a successful operation (e.g., `OrderCreated`, `UserRegistered`).
- Dispatching async tasks to worker queues.
- Decoupling services via event-driven communication.
- Fan-out message distribution to multiple consumers.

## Folder Structure

```
rmq/
└── {feature}/
    ├── {feature}.go        → Constructor & struct (implements outbound port)
    ├── publish.go          → Publish operations
    └── dto/                → Message payload structures (serialize to queue)
```

## Implementation Guide

### 1. Define Outbound Port (in `core/port/outbound/`)

```go
// internal/core/port/outbound/event_publisher.go
package outbound

import "context"

type EventPublisher interface {
    PublishOrderCreated(ctx context.Context, orderID string) error
}
```

### 2. Implement Adapter

```go
// rmq/event/event.go
package event

import (
    "go-simple-template/internal/core/port/outbound"
    amqp "github.com/rabbitmq/amqp091-go"
)

type publisher struct {
    channel  *amqp.Channel
    exchange string
}

func New(ch *amqp.Channel, exchange string) outbound.EventPublisher {
    return &publisher{channel: ch, exchange: exchange}
}
```

```go
// rmq/event/publish.go
package event

import (
    "context"
    "encoding/json"

    amqp "github.com/rabbitmq/amqp091-go"
)

func (p *publisher) PublishOrderCreated(ctx context.Context, orderID string) error {
    body, _ := json.Marshal(map[string]string{"order_id": orderID})
    return p.channel.PublishWithContext(ctx,
        p.exchange,           // exchange
        "order.created",      // routing key
        false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
rmqConn := infrastructure.NewRabbitMQ(cfg)
rmqChannel, _ := rmqConn.Channel()
eventPublisher := rmqEvent.New(rmqChannel, "events-exchange")
orderService := orderApp.New(orderRepo, eventPublisher)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/rabbitmq/amqp091-go` | RabbitMQ AMQP 0.9.1 client |

## Architecture Rules

- Implementations MUST satisfy `core/port/outbound` interfaces.
- Serialize domain events to JSON/protobuf in this layer.
- The domain layer publishes via the port interface — it has no knowledge of RabbitMQ.
