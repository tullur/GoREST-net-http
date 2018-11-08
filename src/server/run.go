package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tullur/rest/src/controller"
	"github.com/tullur/rest/src/pkg"
)

func Run() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	handler := http.NewServeMux()

	handler.HandleFunc("/movie/", pkg.Authantication(pkg.Logger(controller.HandleMovie)))
	handler.HandleFunc("/movies/", pkg.Authantication(pkg.Logger(controller.HandleMovies)))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("Listening on %s\n", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()

	pkg.Graceful(s, 5*time.Second)
}
