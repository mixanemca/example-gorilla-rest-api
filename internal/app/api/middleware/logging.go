// Package middleware for run logics before main program logics
package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"log/slog"
)

// LoggingMiddleware middleware for logging requests
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/swagger/") {
				next.ServeHTTP(w, r)
				return
			}

			body, _ := io.ReadAll(r.Body)
			logger.Info("Request: " + r.Method + " " + r.URL.Path + string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			next.ServeHTTP(w, r)
		})
	}
}
