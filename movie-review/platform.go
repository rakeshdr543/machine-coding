package main

import (
	"fmt"
	"sort"
)

type Platform struct {
	Movies map[string]*Movie
	Users  map[string]*User
}

var platform *Platform

func NewPlatform() *Platform {
	if platform == nil {
		platform = &Platform{Movies: make(map[string]*Movie), Users: make(map[string]*User)}
	}
	return platform
}

func (p *Platform) AddMovie(name string, releaseYear int, genre Genre) error {

	if p.Movies[name] != nil {
		return fmt.Errorf("movie already exists")
	}

	p.Movies[name] = NewMovie(name, releaseYear, genre)
	return nil
}

func (p *Platform) AddUser(username string) error {
	if p.Movies[username] != nil {
		return fmt.Errorf("user already exists")
	}

	p.Users[username] = NewUser(username)
	return nil
}

func (p *Platform) AddMovieReview(movieName string, rating int, reviewerName string) error {
	movie := p.Movies[movieName]
	if movie == nil {
		return fmt.Errorf("movie not found")
	}

	reviewer := p.Users[reviewerName]
	if reviewer == nil {
		return fmt.Errorf("user not found")
	}

	if reviewer.reviews[movieName] != 0 {
		return fmt.Errorf("review already exists")
	}

	reviewer.reviews[movieName] = rating
	reviewer.numberOfReviews++
	if reviewer.numberOfReviews > 3 {
		reviewer.role = Critic
	}

	return nil

}

func (p *Platform) UpdateMovieReview(movieName string, rating int, reviewerName string) error {
	movie := p.Movies[movieName]
	if movie == nil {
		return fmt.Errorf("movie not found")
	}

	reviewer := p.Users[reviewerName]
	if reviewer == nil {
		return fmt.Errorf("user not found")
	}

	if reviewer.reviews[movieName] == 0 {
		return fmt.Errorf("review does not exist")
	}

	reviewer.reviews[movieName] = rating
	return nil
}

func (p *Platform) DeleteMovieReview(movieName string) error {
	movie := p.Movies[movieName]
	if movie == nil {
		return fmt.Errorf("movie not found")
	}

	for _, user := range p.Users {
		if user.reviews[movieName] != 0 {
			delete(user.reviews, movieName)
			user.numberOfReviews--
			if user.numberOfReviews <= 3 {
				user.role = Viewer
			}
		}
	}

	return nil
}

func (p *Platform) AllReviewsOfUser(username string) error {
	user := p.Users[username]
	if user == nil {
		return fmt.Errorf("user not found")
	}

	for movieName, rating := range user.reviews {
		fmt.Printf("Movie: %s, Rating: %d\n", movieName, rating)
	}

	return nil
}

func (p *Platform) TopPlatformMovies(n int) error {
	var Movies []*Movie
	for _, movie := range p.Movies {
		Movies = append(Movies, movie)
	}

	if n > len(Movies) {
		return fmt.Errorf("not enough Movies")
	}

	sort.Slice(Movies, func(i, j int) bool {
		return p.calculateMovieRating(Movies[i].name) > p.calculateMovieRating(Movies[j].name)
	})

	fmt.Printf("Top %d Movies on the platform\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("Movie: %s, Rating: %d\n", Movies[i].name, p.calculateMovieRating(Movies[i].name))
	}

	return nil

}

func (p *Platform) TopGenreMovies(genre Genre, n int) error {
	var Movies []*Movie

	for _, movie := range p.Movies {
		if movie.genre == genre {
			Movies = append(Movies, movie)
		}
	}

	if n > len(Movies) {
		return fmt.Errorf("not enough Movies")
	}

	sort.Slice(Movies, func(i, j int) bool {
		return p.calculateMovieRating(Movies[i].name) > p.calculateMovieRating(Movies[j].name)
	})

	fmt.Printf("Top %d Movies for the genre %d\n", n, genre)

	for i := 0; i < n; i++ {
		fmt.Printf("Movie: %s, Rating: %d\n", Movies[i].name, p.calculateMovieRating(Movies[i].name))
	}

	return nil

}

func (p *Platform) TopMoviesForTheYear(year int, n int) error {
	var Movies []*Movie

	for _, movie := range p.Movies {
		if movie.releaseYear == year {
			Movies = append(Movies, movie)
		}
	}

	if n > len(Movies) {
		return fmt.Errorf("not enough Movies")
	}

	sort.Slice(Movies, func(i, j int) bool {
		return p.calculateMovieRating(Movies[i].name) > p.calculateMovieRating(Movies[j].name)
	})

	fmt.Printf("Top %d Movies for the year %d\n", n, year)

	for i := 0; i < n; i++ {
		fmt.Printf("Movie: %s, Rating: %d\n", Movies[i].name, p.calculateMovieRating(Movies[i].name))
	}

	return nil
}

func (p *Platform) calculateMovieRating(movieName string) int {

	var sum int
	var count int
	for _, user := range p.Users {
		if user.reviews[movieName] != 0 {
			if user.role == Critic {
				sum += user.reviews[movieName] * 2
				count += 2
			} else {
				sum += user.reviews[movieName]
				count++
			}
		}
	}

	if count == 0 {
		return 0
	}

	return sum / count
}
