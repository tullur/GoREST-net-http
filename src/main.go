package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.NewServeMux()

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
