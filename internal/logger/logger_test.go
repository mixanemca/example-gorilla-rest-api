package logger

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
	"testing/slogtest"

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

func TestJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	logger := slog.New(h)
	logger.Debug("Something went wrong")

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Fatal(err)
	}
}
