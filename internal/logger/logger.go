package logger

import (
	"context"
	"log/slog"
	"os"
)

type loggerHandler struct {
	slog.Handler
}

func (h *loggerHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		r.AddAttrs(slog.String("request_id", reqID))
	}

	return h.Handler.Handle(ctx, r)
}

func New(env string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	var handler slog.Handler
	switch env {
	case "production":
		handler = slog.NewJSONHandler(os.Stdout, opts)

	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(&loggerHandler{handler})
}
