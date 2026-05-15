# External API Outbound Adapter

## Purpose

Driven adapter that implements outbound port interfaces by calling **external HTTP/REST APIs**. Used when the application needs to interact with third-party services or other internal microservices over HTTP.

## When to Use

- Integrating with third-party APIs (payment gateways, email providers, SMS services).
- Calling other internal microservices via REST.
- Fetching data from external data sources.
- Webhook delivery.

## Folder Structure

```
extapi/
└── {service_name}/
    ├── {service_name}.go   → Constructor, HTTP client setup, base URL config
    ├── dto/                → Request/response DTOs for the external API
    │   ├── request.go
    │   └── response.go
    ├── {operation}.go      → Individual API call implementations
    └── mapper.go           → Map external DTOs ↔ domain entities
```

## Implementation Guide

### 1. Define Outbound Port (in `core/port/outbound/`)

```go
// internal/core/port/outbound/payment_gateway.go
package outbound

import (
    "context"
    "go-simple-template/internal/core/domain/entity"
)

type PaymentGateway interface {
    Charge(ctx context.Context, payment *entity.Payment) (*entity.PaymentResult, error)
    Refund(ctx context.Context, transactionID string) error
}
```

### 2. Implement Adapter

```go
// extapi/midtrans/midtrans.go
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

```go
// extapi/midtrans/charge.go
package midtrans

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "go-simple-template/internal/core/domain/entity"
    "go-simple-template/internal/adapter/outbound/extapi/midtrans/dto"
    "go-simple-template/internal/pkg/errs"
)

func (c *client) Charge(ctx context.Context, payment *entity.Payment) (*entity.PaymentResult, error) {
    reqBody := dto.ChargeRequest{
        OrderID: payment.OrderID,
        Amount:  payment.Amount,
    }

    body, _ := json.Marshal(reqBody)
    req, _ := http.NewRequestWithContext(ctx, http.MethodPost,
        fmt.Sprintf("%s/v1/charges", c.baseURL), bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, errs.NewInternal(err, "failed to call payment gateway")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errs.NewInternal(nil, "payment gateway returned non-OK status")
    }

    var result dto.ChargeResponse
    json.NewDecoder(resp.Body).Decode(&result)

    return &entity.PaymentResult{
        TransactionID: result.TransactionID,
        Status:        result.Status,
    }, nil
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
httpClient := &http.Client{Timeout: 10 * time.Second}
paymentGateway := midtrans.New(httpClient, cfg.MidtransBaseURL, cfg.MidtransAPIKey)
orderService := orderApp.New(orderRepo, paymentGateway)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `net/http` | Standard HTTP client (built-in) |
| `github.com/go-resty/resty/v2` | Alternative: higher-level REST client |
| `github.com/hashicorp/go-retryablehttp` | HTTP client with automatic retries |

## Architecture Rules

- Implementations MUST satisfy `core/port/outbound` interfaces.
- Map external API DTOs to domain entities in this layer — domain MUST NOT know about external API shapes.
- Handle HTTP errors gracefully and translate to domain errors (`pkg/errs`).
- Use `context.Context` for timeout/cancellation propagation.
- Keep API keys and base URLs as constructor parameters, never hardcode.
