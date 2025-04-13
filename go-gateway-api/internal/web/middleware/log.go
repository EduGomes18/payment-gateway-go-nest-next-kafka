package middleware

import (
	"log"
	"net/http"
	"time"
)

type LogMiddleware struct{}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log request details
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		log.Printf("Headers: %v", r.Header)
		
		// Create a custom response writer to capture status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		// Call the next handler
		next.ServeHTTP(rw, r)
		
		// Log response details
		duration := time.Since(start)
		log.Printf("Response: %d %s (took %v)", rw.statusCode, http.StatusText(rw.statusCode), duration)
	})
}

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
} 