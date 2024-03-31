package main

import (
	"fmt"

	foodkart "github.com/rakeshdr543/machine_coding/foodkart"
)

func main() {

	foodkartApp := foodkart.NewPlatform()

	err := foodkartApp.AddRestaurant(
		"Restaurant1",
		[]foodkart.ItemDetails{
			{Name: "Pizza", Price: 10},
			{Name: "Burger", Price: 5},
		}, 20)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = foodkartApp.AddRestaurant("Restaurant2", []foodkart.ItemDetails{
		{Name: "Pizza", Price: 15},
		{Name: "Burger", Price: 7},
	}, 30)
	if err != nil {
		fmt.Println(err)
		return
	}

	// update menu

	err = foodkartApp.UpdateRestaurantMenu(
		"Restaurant1",
		[]foodkart.ItemDetails{
			{Name: "Pizza", Price: 3},
			{Name: "Burger", Price: 4},
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	// place order
	order, err := foodkartApp.PlaceOrder("rakeshdr", []foodkart.OrderItem{
		{
			Name: "Pizza", Quantity: 6,
		},
		{
			Name: "Burger", Quantity: 2,
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(order)

	foodkartApp.PrintRestaurant()
	foodkartApp.PrintOrder()

	foodkartApp.MarkOrderAsDelivered(order.Id)

	foodkartApp.PrintRestaurant()
	foodkartApp.PrintOrder()
}
