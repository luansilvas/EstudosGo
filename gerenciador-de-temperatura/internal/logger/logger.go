package logger

import (
	"log/slog"
	"os"
)

func New(level string) *slog.Logger {
	var opts slog.HandlerOptions
	switch level {
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "INFO":
		opts.Level = slog.LevelInfo
	default:
		opts.Level = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stdout, &opts)

	return slog.New(handler)
}
