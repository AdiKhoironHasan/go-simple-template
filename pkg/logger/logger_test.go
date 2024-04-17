package logger

import (
	"go-simple-template/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	config.LoadEnv("../../.env")

	logger, err := Init()
	assert.NoError(t, err)

	logger.Debug("test", zap.String("key", "value"))
	logger.Info("test", zap.String("key", "value"))
	logger.Warn("test", zap.String("key", "value"))
	logger.Error("test", zap.String("key", "value"))
	// logger.Fatal("test", zap.String("key", "value"))
}
