package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tullur/rest/src/pkg"
)

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	handler := http.NewServeMux()

	handler.HandleFunc("/", pkg.Logger(index))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Listening on port", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()

	pkg.Graceful(s, 5*time.Second)
}
