package main

type User struct {
	username        string
	role            Role
	reviews         map[string]int
	numberOfReviews int
}

func NewUser(username string) *User {
	return &User{username: username, role: Viewer, reviews: make(map[string]int), numberOfReviews: 0}
}
