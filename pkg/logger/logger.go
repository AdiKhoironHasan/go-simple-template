package logger

import (
	"context"
	"go-simple-template/config"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logKey struct{}

var (
	once   sync.Once
	logCfg zap.Config
	logger *zap.Logger
)

func Init() *zap.Logger {
	once.Do(func() {
		var (
			isDevelopment = config.AppEnv() == "development"
			minLogLevel   = zap.DebugLevel
		)

		if !isDevelopment {
			minLogLevel = zap.InfoLevel
		}

		logCfg = zap.Config{
			Development:      false,
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			Level:            zap.NewAtomicLevelAt(minLogLevel),
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "time",
				CallerKey:      "caller",
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			DisableStacktrace: true,
		}

		zapLog, err := logCfg.Build()
		if err != nil {
			panic(err)
		}

		logger = zapLog

	})

	return logger
}

func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(logKey{}).(*zap.Logger); ok {
		return l
	}

	return nil
}

func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(logKey{}).(*zap.Logger); ok {
		if lp == l {
			return ctx
		}
	}

	return context.WithValue(ctx, logKey{}, l)
}
