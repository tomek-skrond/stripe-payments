package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

type LogEntry struct {
	Timestamp             string `json:"timestamp"`
	Method                string `json:"method"`
	Path                  string `json:"path"`
	StatusCode            int    `json:"status_code"`
	Latency               string `json:"latency"`
	RemoteIP              string `json:"remote_ip"`
	UserAgent             string `json:"user_agent"`
	TLSNegotiatedProtocol string `json:"negotiated_proto"`
	CipherSuite           uint16 `json:"cipher_suite"`
	TLSVersion            string `json:"tls_version"`
}

func JSONLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Capture the response status
		rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(rw, r)

		// Calculate latency
		latency := time.Since(start)

		// Create the log entry
		logEntry := LogEntry{
			Timestamp:  time.Now().Format(time.RFC3339),
			Method:     r.Method,
			Path:       r.URL.Path,
			StatusCode: rw.Status(),
			Latency:    latency.String(),
			RemoteIP:   r.RemoteAddr,
			UserAgent:  r.UserAgent(),
		}

		// Convert log entry to JSON
		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			log.Printf("Could not marshal log entry: %v", err)
			return
		}

		// Output the JSON log
		log.Println(string(logJSON))
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// CORS middleware function
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Change "*" to specific origin if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
