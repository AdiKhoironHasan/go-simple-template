package redis_test

import (
	"go-simple-template/cache/redis"
	"go-simple-template/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.NoError(t, err)

	cfg := config.NewConfig()
	redis := redis.NewRedis(cfg)

	t.Run("success", func(t *testing.T) {
		str, err := redis.Ping()

		assert.NoError(t, err)
		assert.Equal(t, "PONG", str)
	})
}
