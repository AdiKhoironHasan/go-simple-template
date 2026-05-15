# Storage Adapters

Outbound adapters for **file/object storage** services.

## Structure

```
storage/
├── s3/                 # AWS S3 implementations
│   └── <feature>/      # e.g., documents/, media/
│       ├── storage.go          # Constructor + struct
│       └── <method>.go         # e.g., upload.go, download.go
├── gcs/                # Google Cloud Storage
│   └── <feature>/
│       ├── storage.go
│       └── <method>.go
└── README.md
```

## Implementation Guide

```go
package documents

import (
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "go-simple-template/internal/core/port/outbound"
)

type storage struct {
    client *s3.Client
    bucket string
}

func New(client *s3.Client, bucket string) outbound.DocumentStorage {
    return &storage{
        client: client,
        bucket: bucket,
    }
}
```

## Rules

- Storage adapters implement `outbound.XxxStorage` port interfaces
- Pre-signed URL generation lives here
- File validation (size, type) can be shared via `internal/pkg`
- Cloud-specific SDKs are only imported within their sub-package
