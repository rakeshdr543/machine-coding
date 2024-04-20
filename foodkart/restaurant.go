package foodkart

type Restaurant struct {
	Name           string
	menu           map[string]int
	total_capacity int
	used_capacity  int
}
