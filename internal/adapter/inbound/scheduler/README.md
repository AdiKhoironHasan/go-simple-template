# Scheduler (Inbound Adapter)

## Purpose

Driving adapter that triggers application use cases on a **scheduled/cron basis**. Acts as a time-driven entry point, invoking inbound port interfaces at configured intervals.

## When to Use

- Periodic background jobs (e.g., daily report generation, data cleanup).
- Scheduled data synchronization between systems.
- Recurring health checks or monitoring tasks.
- Batch processing at specific intervals.

## Folder Structure

```
scheduler/
├── scheduler.go        → Cron scheduler setup & job registration
└── job/                → Individual job implementations
    └── {feature}/      → Feature-specific scheduled job
```

## Implementation Guide

### 1. Implement a Job

```go
// job/cleanup/job.go
package cleanup

import (
    "context"
    "log/slog"

    "go-simple-template/internal/core/port/inbound"
)

type Job struct {
    service inbound.CleanupService
}

func New(service inbound.CleanupService) *Job {
    return &Job{service: service}
}

func (j *Job) Run() {
    ctx := context.Background()
    if err := j.service.CleanExpiredSessions(ctx); err != nil {
        slog.Error("cleanup job failed", "error", err)
    }
}
```

### 2. Register in Scheduler

```go
// scheduler.go
package scheduler

import "github.com/robfig/cron/v3"

func New(jobs ...JobEntry) *cron.Cron {
    c := cron.New()
    for _, j := range jobs {
        c.AddJob(j.Schedule, j.Job)
    }
    return c
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
cleanupJob := cleanupJob.New(cleanupService)
sched := scheduler.New(
    scheduler.JobEntry{Schedule: "0 2 * * *", Job: cleanupJob}, // daily at 2 AM
)
sched.Start()
defer sched.Stop()
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/robfig/cron/v3` | Cron scheduler |

## Architecture Rules

- Jobs MUST depend on `core/port/inbound` interfaces.
- Jobs are just **triggers** — business logic stays in `app/` services.
- Error handling and logging should be done within the job; the scheduler only manages timing.
