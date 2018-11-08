package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tullur/rest/src/pkg"
)

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", pkg.Logger(index))

	s := http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening on port", s.Addr)
	log.Println(s.ListenAndServe())
}
