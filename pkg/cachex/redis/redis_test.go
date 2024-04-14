package redis_test

import (
	"context"
	"go-simple-template/config"
	"go-simple-template/pkg/cachex/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	config.LoadEnv("../../../.env")
	redis := redis.NewRedis()

	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		str, err := redis.Ping(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "PONG", str)
	})
}
