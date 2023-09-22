package logger

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	l, err := parseLevel("error")
	assert.Nil(t, err)
	assert.Equal(t, slog.LevelError, l)

	l, err = parseLevel("warn")
	assert.Nil(t, err)
	assert.Equal(t, slog.LevelWarn, l)

	l, err = parseLevel("warning")
	assert.Nil(t, err)
	assert.Equal(t, slog.LevelWarn, l)

	l, err = parseLevel("info")
	assert.Nil(t, err)
	assert.Equal(t, slog.LevelInfo, l)

	l, err = parseLevel("debug")
	assert.Nil(t, err)
	assert.Equal(t, slog.LevelDebug, l)

	_, err = parseLevel("invalid")
	assert.Equal(t, "not a valid log level: \"invalid\"", err.Error())
}
