package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// New creates a new logger with the received level and format.
func New(level, format string) (*slog.Logger, error) {
	var logger *slog.Logger

	l, err := parseLevel(level)
	if err != nil {
		return nil, err
	}
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: l}))
	if format == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l}))
	}

	return logger, nil
}

// parseLevel takes a string level and returns the log/slog log level constant.
func parseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(level) {
	case "error":
		return slog.LevelError, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "info":
		return slog.LevelInfo, nil
	case "debug":
		return slog.LevelDebug, nil
	}

	var l slog.Level
	return l, fmt.Errorf("not a valid log level: %q", level)
}
