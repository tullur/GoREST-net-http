package controller

import (
	"fmt"
)

// GetMovies -> show all movies in base
func (s *MovieList) GetMovies() []Movie {
	return s.movies
}

// FindMovieByID -> find by ID
func (s *MovieList) FindMovieByID(id string) *Movie {
	for _, movie := range s.movies {
		if movie.ID == id {
			return &movie
		}
	}

	return nil
}

// AddMovie -> add movie to base
func (s *MovieList) AddMovie(movie Movie) error {
	mv := s.FindMovieByID(movie.ID)
	if mv != nil {
		return fmt.Errorf("Movie with id %s already exists", movie.ID)
	}

	s.movies = append(s.movies, movie)
	return nil
}

// SetMovie -> set movie in base
func (s *MovieList) SetMovie(movie Movie) error {
	for i, mv := range s.movies {
		if mv.ID == movie.ID {
			s.movies[i] = movie
			return nil
		}
	}

	return fmt.Errorf("There is no movie with id %s", movie.ID)
}

// DeleteMovie -> delete movie from base
func (s *MovieList) DeleteMovie(id string) error {
	for i, mv := range s.movies {
		if mv.ID == id {
			s.movies = append(s.movies[:i], s.movies[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("There is no movie with id %s", id)
}
