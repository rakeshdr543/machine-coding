package foodkart

import "errors"

var (
	RestaurantNotFoundError        = errors.New("restaurant not found")
	ErrRestaurantNameAlreadyExists = errors.New("restaurant name already exists")
	OrderNotFoundError             = errors.New("order not found")
)
