package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogRequest is a standard middleware function for doing simple logging of incoming request parameters.
//
// NOTE: Never add the authentication header or unparsed body text to this function or you risk exposing
// user passwords.
func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Host)
		h.ServeHTTP(w, r)
	})
}

// TimeRequest is a standard middleware function for logging the round-trip time to process a request. This
// could be extended to send request times to a monitoring service to detect spikes in processing time.
func TimeRequest(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("Elapsed: %s\n", time.Now().Sub(start))
	})
}
