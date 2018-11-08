package controller

import (
	"fmt"
)

func (s *MovieList) GetMovies() []Movie {
	return s.movies
}

func (s *MovieList) FindMovieByID(id string) *Movie {
	for _, movie := range s.movies {
		if movie.ID == id {
			return &movie
		}
	}

	return nil
}

func (s *MovieList) AddMovie(movie Movie) error {
	mv := s.FindMovieByID(movie.ID)
	if mv != nil {
		return fmt.Errorf("Movie with id %s already exists", movie.ID)
	}

	s.movies = append(s.movies, movie)
	return nil
}

func (s *MovieList) SetMovie(movie Movie) error {
	for i, mv := range s.movies {
		if mv.ID == movie.ID {
			s.movies[i] = movie
			return nil
		}
	}

	return fmt.Errorf("There is no movie with id %s", movie.ID)
}

func (s *MovieList) DeleteMovie(id string) error {
	for i, mv := range s.movies {
		if mv.ID == id {
			s.movies = append(s.movies[:i], s.movies[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("There is no movie with id %s", id)
}
