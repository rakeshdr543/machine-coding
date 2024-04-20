package main

import (
	"fmt"

	foodkart "github.com/rakeshdr543/machine_coding/foodkart"
)

func main() {
	lowerPriceSelectionStrategy := foodkart.LowerPriceSelectionStrategy{}
	foodkartApp := foodkart.NewPlatform(lowerPriceSelectionStrategy)

	// add restaurant
	err := foodkartApp.AddRestaurant("Restaurant1", map[string]int{
		"Pizza":  10,
		"Burger": 5,
	}, 20)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = foodkartApp.AddRestaurant("Restaurant2", map[string]int{
		"Pizza":  5,
		"Burger": 10,
	}, 20)

	if err != nil {
		fmt.Println(err)
		return
	}

	// update menu

	err = foodkartApp.UpdateRestaurantMenu(
		"Restaurant1",
		map[string]int{
			"Pizza":  3,
			"Burger": 5,
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	// place order
	order, err := foodkartApp.PlaceOrder("rakeshdr", map[string]int{
		"Pizza":  2,
		"Burger": 3,
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
