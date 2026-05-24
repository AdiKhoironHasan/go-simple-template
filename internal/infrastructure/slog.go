package infrastructure

import (
	"context"
	"log/slog"
	"os"

	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"
)

// NewSlog initializes the slog logger with the specified log level.
// It sets the default logger to a JSON handler that outputs to stdout.
func NewSlog(logLevel slog.Level) {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	}).WithAttrs([]slog.Attr{
		// add custom attributes if needed
	})

	handler := &slogHandler{
		handler: logHandler,
		level:   logLevel,
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// slogHandler is a custom slog handler that adds request ID to the log context.
// It implements slog.Handler interface.
type slogHandler struct {
	handler slog.Handler
	level   slog.Level
}

// Enabled implements slog.Handler interface to check if the handler is enabled for the given level.
func (s slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= s.level
}

// WithAttrs implements slog.Handler interface to add attributes to the log context.
func (s slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return slogHandler{handler: s.handler.WithAttrs(attrs), level: s.level}
}

// WithGroup implements slog.Handler interface to add a group to the log context.
func (s slogHandler) WithGroup(name string) slog.Handler {
	return slogHandler{handler: s.handler.WithGroup(name), level: s.level}
}

// Handle implements slog.Handler interface to handle the log record.
func (s slogHandler) Handle(ctx context.Context, record slog.Record) error {
	if ctx != nil {
		if requestID, ok := ctx.Value(consts.CtxRequestId).(string); ok && requestID != "" {
			record.AddAttrs(slog.String(consts.CtxRequestId.String(), requestID))
		}
	}

	return s.handler.Handle(ctx, record)
}
