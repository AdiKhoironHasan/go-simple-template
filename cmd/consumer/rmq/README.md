# cmd/consumer/rmq — RabbitMQ Consumer Entrypoint

## Purpose

Composition Root for the **RabbitMQ consumer worker**. Wires infrastructure, message handlers, and the consumer loop together.

## Folder Structure

```
rmq/
└── cmd.go          → DI wiring & consumer startup
```

## Implementation Guide

```go
// cmd/consumer/rmq/cmd.go
package rmq

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

    rmqConn := infrastructure.NewRabbitMQ(ctx)

    // 2. Outbound Adapters
    // orderRepo := orderMongoRepo.New(mongo.Client)

    // 3. Application Services
    // orderSvc := orderService.New(orderRepo)

    // 4. Inbound Adapters (Message Handlers)
    // orderHandler := rmqOrderHandler.New(orderSvc)

    // 5. Consumer Setup
    // consumer := rmqConsumer.New(rmqConn, "order.created.queue", orderHandler)
    // go consumer.Start(ctx)

    _ = rmqConn // placeholder

    // 6. Graceful Shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutting down RabbitMQ consumer...")
}
```

## Usage

```bash
go run cmd/main.go consumer-rmq
```

## Registration in `main.go`

```go
{
    Use:   "consumer-rmq",
    Short: "Start the RabbitMQ consumer worker",
    Run: func(cmd *cobra.Command, args []string) {
        rmq.Start(ctx)
    },
},
```

## Notes

- Each queue/topic should have its own handler, mapped to an inbound port interface.
- Use graceful shutdown to finish processing in-flight messages before exiting.
- Consider running multiple instances for horizontal scaling.
