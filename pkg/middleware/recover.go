package middleware

import (
	"fmt"
	"go-start/pkg/handler"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[ERROR] panic: %v\n%s", err, debug.Stack())
				handler.WriteJSONError(w, fmt.Sprintf("Server error: %v", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
