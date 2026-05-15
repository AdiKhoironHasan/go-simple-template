# gRPC Inbound Adapter

## Purpose

Driving adapter that exposes the application's use cases via **gRPC** (Google Remote Procedure Call). This adapter translates incoming gRPC requests into calls to inbound port interfaces defined in `internal/core/port/inbound/`.

## When to Use

- Inter-service communication in a microservice architecture.
- High-performance, low-latency binary communication.
- Strongly-typed contracts via Protocol Buffers.
- Bi-directional streaming requirements.

## Folder Structure

```
grpc/
├── handler/        → gRPC service implementations (maps to port/inbound interfaces)
├── interceptor/    → gRPC interceptors (logging, auth, recovery — equivalent to HTTP middleware)
├── server/         → gRPC server setup & configuration
└── proto/          → Generated protobuf Go code (from .proto files)
```

> **Note:** The `.proto` source files should live in a top-level `proto/` or `api/proto/` directory. Only the **generated Go code** belongs here.

## Implementation Guide

### 1. Define Protobuf

```protobuf
// api/proto/health/v1/health.proto
syntax = "proto3";
package health.v1;
option go_package = "go-simple-template/internal/adapter/inbound/grpc/proto/healthv1";

service HealthService {
  rpc Check(CheckRequest) returns (CheckResponse);
}
```

### 2. Generate Go Code

```bash
protoc --go_out=. --go-grpc_out=. api/proto/health/v1/health.proto
```

### 3. Implement Handler

```go
// handler/health/handler.go
package health

import (
    "context"
    "go-simple-template/internal/core/port/inbound"
    pb "go-simple-template/internal/adapter/inbound/grpc/proto/healthv1"
)

type Handler struct {
    pb.UnimplementedHealthServiceServer
    service inbound.HealthService
}

func New(service inbound.HealthService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
    result, err := h.service.CheckHealth(ctx)
    if err != nil {
        return nil, err // map to gRPC status codes via interceptor
    }
    return &pb.CheckResponse{Status: result.Status}, nil
}
```

### 4. Wire in Composition Root (`cmd/`)

```go
grpcServer := grpc.NewServer()
healthHandler := grpcHealthHandler.New(healthService)
pb.RegisterHealthServiceServer(grpcServer, healthHandler)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `google.golang.org/grpc` | gRPC framework |
| `google.golang.org/protobuf` | Protocol Buffers runtime |
| `protoc-gen-go` | Protobuf code generator |
| `protoc-gen-go-grpc` | gRPC code generator |

## Architecture Rules

- Handlers MUST depend on `core/port/inbound` interfaces, never on `app/` directly.
- Proto-generated code is treated as **adapter-layer** artifacts.
- Map domain errors (`pkg/errs`) to gRPC status codes in interceptors.
