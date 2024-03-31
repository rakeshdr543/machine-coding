package main

type CartItem struct {
	item     *Item
	quantity int
}

type User struct {
	username string
	address  string
	balance  int
	cart     []*CartItem
}
