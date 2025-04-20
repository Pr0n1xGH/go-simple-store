package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s %s", r.Method, r.URL.Path, r.Header)
		next.ServeHTTP(w, r)
	})
}
