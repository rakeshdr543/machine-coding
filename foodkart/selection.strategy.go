package foodkart

import (
	"fmt"
	"math"
)

type SelectionStrategy interface {
	SelectRestaurant(items map[string]int, restaurants map[string]*Restaurant) (*Restaurant, error)
}

type LowerPriceSelectionStrategy struct {
}

func (l LowerPriceSelectionStrategy) SelectRestaurant(orderItems map[string]int, restaurants map[string]*Restaurant) (*Restaurant, error) {
	lowestPrice := math.MaxUint32
	var lowPriceRestaurant *Restaurant

	for _, restaurant := range restaurants {
		totalPrice := 0
		totalItems := 0

		if restaurant.total_capacity == restaurant.used_capacity {
			continue
		}

		for orderItemName, orderItemQty := range orderItems {
			for itemName, itemPrice := range restaurant.menu {
				if itemName == orderItemName {
					totalPrice += itemPrice * orderItemQty
					totalItems++
				}

			}

			if totalItems == len(orderItems) && totalPrice < lowestPrice {
				lowestPrice = totalPrice
				lowPriceRestaurant = restaurant
			}

		}
	}

	if lowPriceRestaurant == nil {
		return nil, fmt.Errorf(" cannot process by any restaurant")
	}

	return lowPriceRestaurant, nil
}
