package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// HandleMovies -> GET
func HandleMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		HandleGetMovies(w, r)
	}
}

// HandleGetMovies -> movies handler
func HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	movies, err := json.Marshal(MovieBase.GetMovies())
	if err != nil {
		log.Println(err)
	}

	w.Write(movies)
}

// HandleMovie -> Router
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

// HandleGetMovie -> GET handler for movie
func HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	movieID := strings.Replace(r.URL.Path, "/movie/", "", 1)

	movie := MovieBase.FindMovieByID(movieID)
	if movie == nil {
		w.WriteHeader(http.StatusNotFound)

		error, _ := json.Marshal(fmt.Sprintf("Movie with ID: %s not found", movieID))

		w.Write(error)
		return
	}

	w.WriteHeader(http.StatusOK)

	movieIDjson, err := json.Marshal(movie)
	if err != nil {
		log.Println(err)
	}

	w.Write(movieIDjson)
}

// HandleAddMovie -> add handler
func HandleAddMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie Movie
	err := decoder.Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error, _ := json.Marshal(fmt.Sprintf("Bad request. |%v|", err))

		w.Write(error)
		return
	}

	err = MovieBase.AddMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error, _ := json.Marshal(fmt.Sprintf("Bad request. |%v|", err))

		w.Write(error)
		return
	}

	HandleGetMovies(w, r)
}

// HandleUpdateMovie -> update handler
func HandleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieID := strings.Replace(r.URL.Path, "/movie/", "", 1)

	decoder := json.NewDecoder(r.Body)

	var movie Movie

	err := decoder.Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error, _ := json.Marshal(fmt.Sprintf("Bad request. |%v|", err))

		w.Write(error)
		return
	}

	movie.ID = movieID

	err = MovieBase.SetMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		error, _ := json.Marshal(fmt.Sprintf("Bad request. |%v|", err))

		w.Write(error)
		return
	}

	HandleGetMovies(w, r)
}

// HandleDeleteMovie -> delete handler
func HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := strings.Replace(r.URL.Path, "/movie/", "", 1)

	err := MovieBase.DeleteMovie(movieID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		error, _ := json.Marshal(fmt.Sprintf("Bad request. |%v|", err))

		w.Write(error)
		return
	}

	HandleGetMovies(w, r)
}

// HandleMethodIsNotAllowed -> handle not allowed method
func HandleMethodIsNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	message, _ := json.Marshal(fmt.Sprintf("Method %s not allowed", r.Method))
	w.Write(message)
}
