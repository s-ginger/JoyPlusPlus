package logger

import (
	"log/slog"
	"os"
)

var GlobalLogLevel slog.LevelVar

var Global = func() *slog.Logger {
	GlobalLogLevel.Set(slog.LevelInfo)

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &GlobalLogLevel,
	})

	return slog.New(handler)
}()