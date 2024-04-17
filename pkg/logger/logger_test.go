package logger

import (
	"go-simple-template/config"
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	config.LoadEnv("../../.env")

	logger := Init()

	logger.Debug("test", zap.String("key", "value"))
	logger.Info("test", zap.String("key", "value"))
	logger.Warn("test", zap.String("key", "value"))
	logger.Error("test", zap.String("key", "value"))
}
