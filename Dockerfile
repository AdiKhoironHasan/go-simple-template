# Build Stage
FROM golang:1.20.5-alpine AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download

COPY . .

RUN go build -o main

# Deploy Stage
FROM alpine:3.18.4 AS DeployStage

# RUN apt-get update && apt-get install -y --no-install-recommends \
#     ca-certificates \
#     && rm -rf /var/lib/apt/lists/*

RUN apk add curl tzdata ca-certificates && \
    rm -rf /var/cache/apk/*

ARG GCS_PATH
ENV TZ=Asia/Jakarta
# ENV DEBIAN_FRONTEND noninteractive

WORKDIR /app

COPY --from=builder ./app/main .
COPY --from=builder ./app/.env .

CMD ["./main"]
