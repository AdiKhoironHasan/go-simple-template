package infrastructure

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"

	"go-simple-template/internal/pkg/config"
	"go-simple-template/internal/pkg/consts"

	red "github.com/redis/go-redis/v9"
)

type redis struct {
	Client        *red.Client
	connectionURL string
}

func NewRedis(ctx context.Context) *redis {
	return &redis{
		connectionURL: fmt.Sprintf("%s:%d", config.RedisHost(), config.RedisPort()),
	}
}

// Create Redis client with TLS enabled
func newTlsClient(opts red.Options) *red.Client {
	// Enable TLS if it's not already set
	if opts.TLSConfig == nil {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // Set to false in production to validate certificates
		}
	}

	return red.NewClient(&opts)
}

func (a *redis) Connect(ctx context.Context) {

	var client *red.Client
	rdbOpts := &red.Options{
		Addr:     a.connectionURL,
		Username: config.RedisUsername(),
		Password: config.RedisPassword(),
		DB:       config.RedisDB(),
	}

	// Choose whether to use TLS or not
	if config.RedisTLS() {
		client = newTlsClient(*rdbOpts)
	} else {
		client = red.NewClient(rdbOpts)
	}

	slog.InfoContext(ctx, "Redis connected", slog.String("connected to", a.connectionURL))

	a.Client = client
}

func (a *redis) Disconnect(ctx context.Context) {
	err := a.Client.Close()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to disconnect to Redis", slog.String("connection", a.connectionURL), slog.String(consts.Error, err.Error()))
	}
}
