package redis_test

import (
	"context"
	"go-simple-template/config"
	"go-simple-template/pkg/cachex/redis"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.NoError(t, err)

	ctx := context.Background()
	cfg := config.NewConfig()
	redis := redis.NewRedis(cfg)

	t.Run("success", func(t *testing.T) {
		str, err := redis.Ping(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "PONG", str)
	})
}
