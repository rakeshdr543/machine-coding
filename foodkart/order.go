package foodkart

type OrderItem struct {
	Name     string
	Quantity int
}

type OrderDetails struct {
	restaurant *Restaurant
	items      []OrderItem
}

type Order struct {
	Id            string
	userName      string
	order_details []OrderDetails
	delivered     bool
}
