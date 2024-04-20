package foodkart

import (
	"fmt"
	"time"
)

type Platform struct {
	restaurants       map[string]*Restaurant
	orders            map[string]*Order
	selectionStrategy SelectionStrategy
}

func NewPlatform(selectionStrategy SelectionStrategy) *Platform {
	return &Platform{restaurants: make(map[string]*Restaurant), orders: make(map[string]*Order), selectionStrategy: selectionStrategy}
}

func (p *Platform) AddRestaurant(Name string, menu map[string]int, capacity int) error {
	if p.restaurants[Name] != nil {
		return ErrRestaurantNameAlreadyExists
	}

	p.restaurants[Name] = &Restaurant{
		Name:           Name,
		menu:           menu,
		total_capacity: capacity,
		used_capacity:  0,
	}

	return nil
}

func (p *Platform) UpdateRestaurantMenu(Name string, menu map[string]int) error {
	restaurant := p.restaurants[Name]
	if restaurant == nil {
		return RestaurantNotFoundError
	}

	restaurant.menu = menu
	return nil
}

func (p *Platform) PlaceOrder(userName string, items map[string]int) (Order, error) {

	restaurant, err := p.selectionStrategy.SelectRestaurant(items, p.restaurants)

	if err != nil {
		return Order{}, err
	}

	orderDetails :=
		OrderDetails{
			restaurant: restaurant,
			items:      items,
		}

	newOrder := &Order{
		Id:            fmt.Sprintf("order_%d", time.Now().Unix()),
		userName:      userName,
		delivered:     false,
		order_details: []OrderDetails{orderDetails},
	}

	restaurant.used_capacity += 1
	p.orders[newOrder.Id] = newOrder
	return *newOrder, nil
}

func (p *Platform) MarkOrderAsDelivered(Id string) error {
	order := p.orders[Id]

	if order == nil {
		return OrderNotFoundError
	}

	order.delivered = true

	for _, orderData := range order.order_details {
		restaurant := orderData.restaurant
		restaurant.used_capacity -= 1
	}

	return nil
}

func (p *Platform) PrintRestaurant() {
	fmt.Println("All Restaurants details")

	for _, restaurant := range p.restaurants {
		fmt.Println(restaurant)
	}
}

func (p *Platform) PrintOrder() {
	fmt.Println("All Orders")

	for _, order := range p.orders {
		fmt.Println(order)
	}
}
