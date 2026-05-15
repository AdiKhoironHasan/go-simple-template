# RabbitMQ Consumer (Inbound Adapter)

## Purpose

Driving adapter that consumes messages from **RabbitMQ** queues and translates them into calls to inbound port interfaces. This enables event-driven processing where the application reacts to messages published by other services.

## When to Use

- Asynchronous task processing (e.g., email sending, report generation).
- Event-driven architecture where services communicate via message queues.
- Work queue patterns with multiple consumers for load distribution.
- Retry/dead-letter queue patterns for fault tolerance.

## Folder Structure

```
rmq/
├── consumer.go         → RabbitMQ consumer setup (connection, channel, queue binding)
├── handler/            → Message handlers (maps queue messages to port/inbound calls)
│   └── {feature}/      → Feature-specific message handler
└── dto/                → Message payload structures (deserialize from queue)
```

## Implementation Guide

### 1. Define Message DTO

```go
// dto/order_created.go
package dto

type OrderCreatedMessage struct {
    OrderID   string `json:"order_id"`
    UserID    string `json:"user_id"`
    Amount    int64  `json:"amount"`
    CreatedAt string `json:"created_at"`
}
```

### 2. Implement Message Handler

```go
// handler/order/handler.go
package order

import (
    "context"
    "encoding/json"

    "go-simple-template/internal/adapter/inbound/consumer/rmq/dto"
    "go-simple-template/internal/core/port/inbound"

    amqp "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
    service inbound.OrderService
}

func New(service inbound.OrderService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, msg amqp.Delivery) error {
    var payload dto.OrderCreatedMessage
    if err := json.Unmarshal(msg.Body, &payload); err != nil {
        msg.Nack(false, false) // reject, don't requeue
        return err
    }

    err := h.service.ProcessOrder(ctx, payload.OrderID)
    if err != nil {
        msg.Nack(false, true) // reject, requeue for retry
        return err
    }

    msg.Ack(false)
    return nil
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
rmqConn := infrastructure.NewRabbitMQ(cfg)
orderHandler := rmqOrderHandler.New(orderService)
consumer := rmqConsumer.New(rmqConn, "order.created.queue", orderHandler)
go consumer.Start(ctx)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/rabbitmq/amqp091-go` | RabbitMQ AMQP 0.9.1 client |

## Architecture Rules

- Handlers MUST depend on `core/port/inbound` interfaces.
- Message deserialization (DTO → domain entity) happens in this layer.
- NEVER import `app/` directly — always go through the port interface.
- Acknowledge/reject messages based on processing result.
