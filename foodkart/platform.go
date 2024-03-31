package foodkart

import (
	"fmt"
	"time"
)

type Platform struct {
	restaurants map[string]*Restaurant
	orders      map[string]*Order
}

func NewPlatform() *Platform {
	return &Platform{restaurants: make(map[string]*Restaurant), orders: make(map[string]*Order)}
}

func (p *Platform) AddRestaurant(Name string, menu []ItemDetails, capacity int) error {
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

func (p *Platform) UpdateRestaurantMenu(Name string, menu []ItemDetails) error {
	restaurant := p.restaurants[Name]
	if restaurant == nil {
		return RestaurantNotFoundError
	}

	restaurant.menu = menu
	return nil
}

type RestaurantOrder struct {
	items []OrderItem
}

func (p *Platform) PlaceOrder(userName string, items []OrderItem) (Order, error) {

	restaurantOrdersMap := make(map[*Restaurant]RestaurantOrder)
	orderDetails := make([]OrderDetails, 0)

	for _, item := range items {
		restaurant, err := p.GetRestaurantOfferingLowerPrice(item)
		if err != nil {
			return Order{}, err
		}
		if _, ok := restaurantOrdersMap[restaurant]; !ok {
			restaurantOrdersMap[restaurant] = RestaurantOrder{
				items: make([]OrderItem, 0),
			}
		}

		restaurantOrder := restaurantOrdersMap[restaurant]
		restaurantOrder.items = append(restaurantOrder.items, item)
		restaurantOrdersMap[restaurant] = restaurantOrder
	}

	for restaurant, orderData := range restaurantOrdersMap {
		orderDetails = append(orderDetails,
			OrderDetails{
				restaurant: restaurant,
				items:      orderData.items,
			})

		restaurant.used_capacity += len(orderData.items)
	}

	newOrder := &Order{
		Id:            fmt.Sprintf("order_%d", time.Now().Unix()),
		userName:      userName,
		delivered:     false,
		order_details: orderDetails,
	}
	p.orders[newOrder.Id] = newOrder
	return *newOrder, nil
}

func (p *Platform) GetRestaurantOfferingLowerPrice(orderItem OrderItem) (*Restaurant, error) {
	lowestPrice := HIGHEST_ITEM_Price
	var lowPriceRestaurant *Restaurant

	for _, restaurant := range p.restaurants {
		for _, item := range restaurant.menu {
			if item.Name == orderItem.Name {
				if item.Price < lowestPrice {
					lowestPrice = item.Price
					lowPriceRestaurant = restaurant
				}

			}
		}
	}

	if lowPriceRestaurant == nil {
		return nil, fmt.Errorf(" cannot find restaurant with Name %s", orderItem.Name)
	}

	return lowPriceRestaurant, nil
}

func (p *Platform) MarkOrderAsDelivered(Id string) error {
	order := p.orders[Id]

	if order == nil {
		return OrderNotFoundError
	}

	order.delivered = true

	for _, orderData := range order.order_details {
		restaurant := orderData.restaurant
		restaurant.used_capacity -= len(orderData.items)
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
