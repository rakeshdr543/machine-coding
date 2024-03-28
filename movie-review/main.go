package main

func main() {
	platform := NewPlatform()

	platform.AddMovie("The Shawshank Redemption", 1994, Drama)
	platform.AddMovie("The Godfather", 1972, Drama)
	platform.AddMovie("The Dark Knight", 2008, Action)
	platform.AddMovie("The Godfather: Part II", 1974, Drama)
	platform.AddMovie("The Lord of the Rings: The Return of the King", 2003, Fantasy)
	platform.AddMovie("Pulp Fiction", 1994, Crime)
	platform.AddMovie("Schindler's List", 1993, Biography)
	platform.AddMovie("The Lord of the Rings: The Fellowship of the Ring", 2001, Fantasy)

	platform.AddUser("Alice")
	platform.AddUser("Bob")
	platform.AddUser("Charlie")
	platform.AddUser("David")

	platform.AddMovieReview("The Shawshank Redemption", 5, "Alice")
	platform.AddMovieReview("The Shawshank Redemption", 4, "Bob")
	platform.AddMovieReview("The Shawshank Redemption", 3, "Charlie")
	platform.AddMovieReview("The Shawshank Redemption", 2, "David")

	platform.AddMovieReview("The Godfather", 5, "Alice")
	platform.AddMovieReview("The Godfather", 4, "Bob")
	platform.AddMovieReview("The Godfather", 3, "Charlie")
	platform.AddMovieReview("The Godfather", 2, "David")

	platform.AddMovieReview("The Dark Knight", 5, "Alice")
	platform.AddMovieReview("The Dark Knight", 4, "Bob")

	platform.TopPlatformMovies(3)
	platform.AllReviewsOfUser("Alice")
	platform.TopGenreMovies(Drama, 3)
	platform.TopMoviesForTheYear(2008, 3)
}
