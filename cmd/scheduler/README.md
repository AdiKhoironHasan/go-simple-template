# cmd/scheduler — Scheduler Entrypoint

## Purpose

Composition Root for the **cron scheduler**. Wires infrastructure, scheduled jobs, and the cron engine together.

## Folder Structure

```
scheduler/
└── cmd.go          → DI wiring & scheduler startup
```

## Implementation Guide

```go
// cmd/scheduler/cmd.go
package scheduler

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "go-simple-template/internal/infrastructure"

    "github.com/robfig/cron/v3"
)

func Start(ctx context.Context) {
    // 1. Infrastructure
    mongo := infrastructure.NewMongoDB(ctx)
    _ = mongo.Connect(ctx)

    redis := infrastructure.NewRedis(ctx)
    redis.Connect(ctx)

    // 2. Outbound Adapters
    // cleanupRepo := cleanupMongoRepo.New(mongo.Client)

    // 3. Application Services
    // cleanupSvc := cleanupService.New(cleanupRepo)

    // 4. Jobs
    // cleanupJob := cleanupJob.New(cleanupSvc)
    // reportJob  := reportJob.New(reportSvc)

    // 5. Cron Setup
    c := cron.New()
    // c.AddJob("0 2 * * *", cleanupJob)   // daily at 2 AM
    // c.AddJob("0 */6 * * *", reportJob)  // every 6 hours
    c.Start()

    log.Println("scheduler started")

    // 6. Graceful Shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutting down scheduler...")
    c.Stop()
}
```

## Usage

```bash
go run cmd/main.go scheduler
```

## Registration in `main.go`

```go
{
    Use:   "scheduler",
    Short: "Start the cron scheduler",
    Run: func(cmd *cobra.Command, args []string) {
        scheduler.Start(ctx)
    },
},
```

## Notes

- Each job should implement the `cron.Job` interface (just a `Run()` method).
- Jobs call inbound port interfaces — business logic stays in `app/` services.
- Use `cron.WithSeconds()` option if you need second-level precision.
- Consider using distributed locks (e.g., Redis-based) if running multiple scheduler instances.
