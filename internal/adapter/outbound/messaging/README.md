# Messaging Adapters

Outbound adapters for **event/message publishing** to message brokers.

## Structure

```
messaging/
├── kafka/              # Kafka producers
│   └── <topic>/        # e.g., order_events/, user_events/
│       ├── publisher.go        # Constructor + struct
│       └── <method>.go         # e.g., publish_order_created.go
├── rabbitmq/           # RabbitMQ publishers
│   └── <exchange>/     # e.g., notifications/, audit/
│       ├── publisher.go
│       └── <method>.go
└── README.md
```

## Implementation Guide

```go
package orderevents

import (
    "github.com/segmentio/kafka-go"
    "go-simple-template/internal/core/port/outbound"
)

type publisher struct {
    writer *kafka.Writer
}

func New(writer *kafka.Writer) outbound.OrderEventPublisher {
    return &publisher{writer: writer}
}
```

## Rules

- Publishers implement `outbound.XxxPublisher` or `outbound.XxxEventPublisher` interfaces
- Message serialization (JSON, Protobuf) happens here, not in domain
- Retry & dead-letter logic lives here or in infrastructure
- One publisher per topic/exchange group
