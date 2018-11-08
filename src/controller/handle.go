package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		HandleGetMovies(w, r)
	}
}

func HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	movies, err := json.Marshal(MovieBase.GetMovies())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(movies)
}

func HandleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		HandleGetMovie(w, r)
	} else if r.Method == http.MethodPost {
		HandleAddMovie(w, r)
	} else if r.Method == http.MethodPut {
		HandleUpdateMovie(w, r)
	} else if r.Method == http.MethodDelete {
		HandleDeleteMovie(w, r)
	} else {
		HandleMethodIsNotAllowed(w, r)
	}
}

// TODO: CRUD Handlers
func HandleGetMovie(w http.ResponseWriter, r *http.Request) {

}

func HandleAddMovie(w http.ResponseWriter, r *http.Request) {

}

func HandleUpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {

}

func HandleMethodIsNotAllowed(w http.ResponseWriter, r *http.Request) {

}

// TODO: Movie Handlers & Create Methods GET/Find/Set/Delete
func (s MovieList) GetMovies() []Movie {
	return s.movies
}
