package main

type Movie struct {
	name        string
	releaseYear int
	genre       Genre
}

func NewMovie(name string, releaseYear int, genre Genre) *Movie {
	return &Movie{name: name, releaseYear: releaseYear, genre: genre}
}
