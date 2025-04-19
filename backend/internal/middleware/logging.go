package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		log.Printf("➡️  %s %s | IP: %s | UA: %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			r.UserAgent(),
		)

		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		log.Printf("✅ %s %s | %d %s | Time: %v",
			r.Method,
			r.URL.Path,
			ww.statusCode,
			http.StatusText(ww.statusCode),
			duration,
		)
	})
}

// Helper to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
