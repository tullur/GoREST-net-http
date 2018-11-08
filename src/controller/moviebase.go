package controller

type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Director    string `json:"director,omitempty"`
	Genre       string `json:"genre,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
}

type MovieList struct {
	movies []Movie
}

var MovieBase = MovieList{
	movies: make([]Movie, 0),
}
