package foodkart

type OrderDetails struct {
	restaurant *Restaurant
	items      map[string]int
}

type Order struct {
	Id            string
	userName      string
	order_details []OrderDetails
	delivered     bool
}
