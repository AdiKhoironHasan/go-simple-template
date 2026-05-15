# cmd/api/grpc — gRPC Server Entrypoint

## Purpose

Composition Root for the **gRPC API server**. Wires infrastructure, outbound adapters, application services, and gRPC handlers together.

## Folder Structure

```
grpc/
└── cmd.go          → DI wiring & gRPC server startup
```

## Implementation Guide

```go
// cmd/api/grpc/cmd.go
package grpc

import (
    "context"
    "log"
    "net"

    "go-simple-template/internal/infrastructure"

    "google.golang.org/grpc"
)

func Start(ctx context.Context) {
    // 1. Infrastructure
    mongo := infrastructure.NewMongoDB(ctx)
    _ = mongo.Connect(ctx)

    redis := infrastructure.NewRedis(ctx)
    redis.Connect(ctx)

    // 2. Outbound Adapters
    // healthRepo := healthMongoRepo.New(mongo.Client)

    // 3. Application Services
    // healthSvc := healthService.New(healthRepo, ...)

    // 4. gRPC Handlers
    // healthHandler := grpcHealthHandler.New(healthSvc)

    // 5. Server Setup
    server := grpc.NewServer(
        // grpc.UnaryInterceptor(...),
    )
    // pb.RegisterHealthServiceServer(server, healthHandler)

    // 6. Run
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    log.Println("gRPC server listening on :50051")
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve gRPC: %v", err)
    }
}
```

## Usage

```bash
go run cmd/main.go grpc
```

## Registration in `main.go`

```go
{
    Use:   "grpc",
    Short: "Start the gRPC API server",
    Run: func(cmd *cobra.Command, args []string) {
        grpc.Start(ctx)
    },
},
```
