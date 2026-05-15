# cmd/consumer/kafka — Kafka Consumer Entrypoint

## Purpose

Composition Root for the **Kafka consumer worker**. Wires infrastructure, message handlers, and the consumer group together.

## Folder Structure

```
kafka/
└── cmd.go          → DI wiring & consumer startup
```

## Implementation Guide

```go
// cmd/consumer/kafka/cmd.go
package kafka

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "go-simple-template/internal/infrastructure"
)

func Start(ctx context.Context) {
    // 1. Infrastructure
    mongo := infrastructure.NewMongoDB(ctx)
    _ = mongo.Connect(ctx)

    kafkaReader := infrastructure.NewKafkaReader(ctx, "analytics-topic", "group-1")

    // 2. Outbound Adapters
    // analyticsRepo := analyticsMongoRepo.New(mongo.Client)

    // 3. Application Services
    // analyticsSvc := analyticsService.New(analyticsRepo)

    // 4. Inbound Adapters (Message Handlers)
    // analyticsHandler := kafkaAnalyticsHandler.New(analyticsSvc)

    // 5. Consumer Setup
    // consumer := kafkaConsumer.New(kafkaReader, analyticsHandler)
    // go consumer.Start(ctx)

    _ = kafkaReader // placeholder

    // 6. Graceful Shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutting down Kafka consumer...")
}
```

## Usage

```bash
go run cmd/main.go consumer-kafka
```

## Registration in `main.go`

```go
{
    Use:   "consumer-kafka",
    Short: "Start the Kafka consumer worker",
    Run: func(cmd *cobra.Command, args []string) {
        kafka.Start(ctx)
    },
},
```

## Notes

- Use consumer groups for parallel processing across partitions.
- Configure auto-commit vs manual commit based on delivery guarantees needed.
- Consider separate entrypoints per topic group for independent scaling.
