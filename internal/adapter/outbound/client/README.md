# Client Adapters

Outbound adapters for **external/3rd-party service** integrations.

## Structure

```
client/
├── rest/               # REST API clients
│   └── <service>/      # e.g., midtrans/, xendit/, sendgrid/
│       ├── client.go           # Constructor + struct + http.Client
│       └── <method>.go         # One file per API call
├── grpc/               # gRPC clients
│   └── <service>/      # e.g., payment/, notification/
│       ├── client.go
│       └── <method>.go
└── README.md
```

## Implementation Guide

```go
package midtrans

import (
    "net/http"
    "go-simple-template/internal/core/port/outbound"
)

type client struct {
    httpClient *http.Client
    baseURL    string
    apiKey     string
}

func New(httpClient *http.Client, baseURL, apiKey string) outbound.PaymentGateway {
    return &client{
        httpClient: httpClient,
        baseURL:    baseURL,
        apiKey:     apiKey,
    }
}
```

## Rules

- Each external service = one sub-package
- HTTP client is injected, not created internally (testability)
- API-specific DTOs live inside the client package (not in domain)
- Map external DTOs to domain entities before returning
