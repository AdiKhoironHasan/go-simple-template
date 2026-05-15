# cmd/ — Entrypoints & Composition Root

## Purpose

This directory contains all **application entrypoints**. Each subdirectory is a **Composition Root** — the single place where all dependencies (infrastructure, adapters, services) are wired together using explicit dependency injection.

The `main.go` at the root uses **Cobra CLI** to register each entrypoint as a subcommand.

## Architecture Role

```
cmd/ is the OUTERMOST layer — it imports everything, but nothing imports cmd/.
```

All dependency injection happens here. This is the **only place** where concrete implementations are bound to port interfaces.

## Folder Structure

```
cmd/
├── main.go                   → Cobra CLI bootstrap (registers all subcommands)
│
├── api/                      → API server entrypoints
│   ├── rest/                 → REST API server (Echo)
│   │   └── cmd.go            → DI wiring: infra → adapters → services → server
│   └── grpc/                 → gRPC API server
│       └── cmd.go            → DI wiring: infra → adapters → services → server
│
├── consumer/                 → Message consumer entrypoints
│   ├── rmq/                  → RabbitMQ consumer worker
│   │   └── cmd.go            → DI wiring: infra → handlers → consumer loop
│   └── kafka/                → Kafka consumer worker
│       └── cmd.go            → DI wiring: infra → handlers → consumer loop
│
└── scheduler/                → Scheduled jobs entrypoint
    └── cmd.go                → DI wiring: infra → jobs → cron scheduler
```

## How It Works

### Cobra Subcommands

Each entrypoint is registered as a Cobra subcommand in `main.go`:

```go
// main.go
cmd := []*cobra.Command{
    {
        Use:   "rest",
        Short: "Start the REST API server",
        Run: func(cmd *cobra.Command, args []string) {
            rest.Start(ctx)
        },
    },
    {
        Use:   "grpc",
        Short: "Start the gRPC API server",
        Run: func(cmd *cobra.Command, args []string) {
            grpc.Start(ctx)
        },
    },
    {
        Use:   "consumer-rmq",
        Short: "Start the RabbitMQ consumer worker",
        Run: func(cmd *cobra.Command, args []string) {
            rmq.Start(ctx)
        },
    },
    {
        Use:   "consumer-kafka",
        Short: "Start the Kafka consumer worker",
        Run: func(cmd *cobra.Command, args []string) {
            kafka.Start(ctx)
        },
    },
    {
        Use:   "scheduler",
        Short: "Start the cron scheduler",
        Run: func(cmd *cobra.Command, args []string) {
            scheduler.Start(ctx)
        },
    },
}
```

### Usage

```bash
# Start REST server
go run cmd/main.go rest

# Start gRPC server
go run cmd/main.go grpc

# Start RabbitMQ consumer
go run cmd/main.go consumer-rmq

# Start Kafka consumer
go run cmd/main.go consumer-kafka

# Start scheduler
go run cmd/main.go scheduler
```

## Composition Root Pattern

Each `cmd.go` follows the same DI wiring pattern:

```
1. Infrastructure  → Create DB clients, message brokers, caches
2. Outbound        → Bind repositories/publishers to port interfaces
3. Application     → Create services with injected dependencies
4. Inbound         → Create handlers/consumers with injected services
5. Run             → Start server/consumer/scheduler
```

### Example: `cmd/api/rest/cmd.go` (existing)

```go
func Start(ctx context.Context) {
    // 1. Infrastructure
    mongo := infrastructure.NewMongoDB(ctx)
    redis := infrastructure.NewRedis(ctx)

    // 2. Outbound Adapters
    healthRepo := healthRepo.New(mongo.Client)
    healthCache := healthCache.NewCache(redis.Client)

    // 3. Application Services
    healthSvc := healthService.New(healthRepo, healthCache)

    // 4. Inbound Adapter (REST)
    deps := &router.Dependencies{HealthService: healthSvc}
    srv := server.New(ctx, deps)

    // 5. Run
    srv.Run(ctx)
}
```

## Rules

- Each `cmd.go` exposes a single `Start(ctx context.Context)` function.
- **NO business logic** in this directory — only wiring.
- **NO shared state** between entrypoints — each is self-contained.
- Import direction: `cmd/ → adapter/ → app/ → core/` and `cmd/ → infrastructure/`.
