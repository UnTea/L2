package api

import (
	"log"
	"net/http"
	"time"
)

// Logging is a function that handling incoming http handlers
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, req)
		log.Printf("method: %s  URI: %s  time: %s", req.Method, req.RequestURI, time.Since(start))
	})
}
