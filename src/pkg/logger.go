package pkg

import (
	"log"
	"net/http"
)

// Logger - function which log handlers.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Server [net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
