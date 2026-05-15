# Elasticsearch Outbound Adapter

## Purpose

Driven adapter that implements outbound port interfaces using **Elasticsearch** as a search/indexing engine. Used for full-text search, log aggregation, and analytics queries.

## When to Use

- Full-text search across domain entities (products, articles, users).
- Log and event aggregation for observability.
- Analytics dashboards with aggregation queries.
- Autocomplete and suggestion features.

## Folder Structure

```
elasticsearch/
└── {feature}/
    ├── {feature}.go        → Constructor & struct (implements outbound port)
    ├── index.go            → Index/create document operations
    ├── search.go           → Search/query operations
    └── mapping.go          → Elasticsearch index mapping definitions
```

## Implementation Guide

### 1. Define Outbound Port (in `core/port/outbound/`)

```go
// internal/core/port/outbound/product_search.go
package outbound

import (
    "context"
    "go-simple-template/internal/core/domain/entity"
)

type ProductSearch interface {
    Index(ctx context.Context, product *entity.Product) error
    Search(ctx context.Context, query string, page, size int) ([]*entity.Product, int64, error)
}
```

### 2. Implement Adapter

```go
// elasticsearch/product/product.go
package product

import (
    "go-simple-template/internal/core/port/outbound"
    "github.com/elastic/go-elasticsearch/v8"
)

type searchRepo struct {
    client    *elasticsearch.Client
    indexName string
}

func New(client *elasticsearch.Client, indexName string) outbound.ProductSearch {
    return &searchRepo{client: client, indexName: indexName}
}
```

```go
// elasticsearch/product/search.go
package product

import (
    "context"
    "encoding/json"
    "strings"

    "go-simple-template/internal/core/domain/entity"
)

func (r *searchRepo) Search(ctx context.Context, query string, page, size int) ([]*entity.Product, int64, error) {
    body := map[string]interface{}{
        "query": map[string]interface{}{
            "multi_match": map[string]interface{}{
                "query":  query,
                "fields": []string{"name^3", "description"},
            },
        },
        "from": (page - 1) * size,
        "size": size,
    }

    bodyJSON, _ := json.Marshal(body)
    res, err := r.client.Search(
        r.client.Search.WithContext(ctx),
        r.client.Search.WithIndex(r.indexName),
        r.client.Search.WithBody(strings.NewReader(string(bodyJSON))),
    )
    if err != nil {
        return nil, 0, err
    }
    defer res.Body.Close()

    // Parse response and map to domain entities...
    return products, total, nil
}
```

### 3. Wire in Composition Root (`cmd/`)

```go
esClient := infrastructure.NewElasticsearch(cfg)
productSearch := esProduct.New(esClient, "products")
productService := productApp.New(productRepo, productSearch)
```

## Key Dependencies

| Package | Purpose |
|---|---|
| `github.com/elastic/go-elasticsearch/v8` | Official Elasticsearch client |
| `github.com/olivere/elastic/v7` | Alternative: community client (v7) |

## Infrastructure Setup

The corresponding client setup belongs in `internal/infrastructure/elasticsearch.go`.

## Architecture Rules

- Implementations MUST satisfy `core/port/outbound` interfaces.
- Return domain entities, NOT Elasticsearch-specific models.
- Index mappings and queries live in this adapter, not in `core/`.
- Handle Elasticsearch errors and map to domain errors (`pkg/errs`).
