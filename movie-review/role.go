package main

type Role int

const (
	Viewer Role = iota
	Critic
	Expert
)

var RoleWeighage = map[Role]int{
	Viewer: 1,
	Critic: 2,
	Expert: 3,
}
