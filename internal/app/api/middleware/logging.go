// Package middleware for run logics before main program logics
package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"log/slog"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{w, http.StatusOK}
}

func (w *ResponseWriterWrapper) Status() int {
	return w.status
}

// LoggingMiddleware middleware for logging successful requests
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/swagger/") {
				next.ServeHTTP(w, r)
				return
			}

			// create a ResponseWriter Wrapper that intercepts the response status
			wrapper := NewResponseWriterWrapper(w)
			next.ServeHTTP(wrapper, r)
			statusCode := wrapper.Status()

			// skip bad request status
			if statusCode != http.StatusBadRequest {
				body, _ := io.ReadAll(r.Body)
				logger.Info("Request: " + r.Method + " " + r.URL.Path + string(body))
				r.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		})
	}
}
