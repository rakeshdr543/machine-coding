package foodkart

type ItemDetails struct {
	Name  string
	Price int
}

type Restaurant struct {
	Name           string
	menu           []ItemDetails
	total_capacity int
	used_capacity  int
}
