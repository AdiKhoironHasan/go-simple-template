# Stage 1: Build the Go application
FROM golang:1.26-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

# Copy dependency files first (cache layer)
COPY go.mod go.sum ./
COPY vendor/ vendor/

# Copy source code
COPY . .

# Build with vendor
RUN go build -mod=vendor -o go-simple-app ./cmd

# Stage 2: Minimal runtime image
FROM alpine:3.22

RUN apk --no-cache add \
    ca-certificates \
    curl \
    tzdata

ARG APP_VERSION

ENV APP_VERSION=$APP_VERSION
ENV TZ=Asia/Jakarta
ENV APP_PORT=80

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/go-simple-app .

EXPOSE 80

CMD ["./go-simple-app", "rest"]
