package main

import (
	"errors"
	"fmt"
)

type ItemIdentifier struct {
	category Category
	brand    Brand
}

type Platform struct {
	users map[string]*User
	items map[ItemIdentifier]*Item
}

var platform *Platform

func NewPlatform() *Platform {
	if platform != nil {
		return platform
	}
	return &Platform{
		users: make(map[string]*User),
		items: make(map[ItemIdentifier]*Item),
	}
}

func (p *Platform) AddUser(username string, address string, balance int) error {
	if p.users[username] != nil {
		return errors.New("User already exists")
	}
	newUser := &User{
		username: username,
		address:  address,
		balance:  balance,
	}

	p.users[username] = newUser

	fmt.Println("User added successfully")
	return nil
}

func (p *Platform) AddItem(brand Brand, category Category, price int) error {
	itemIdentifier := ItemIdentifier{
		category: category,
		brand:    brand,
	}

	if p.items[itemIdentifier] != nil {
		return errors.New("Item already exist")
	}
	p.items[itemIdentifier] = &Item{
		category: category,
		brand:    brand,
		price:    price,
	}
	fmt.Println("Item added successfully")
	return nil
}

func (p *Platform) AddItemInventory(brand Brand, category Category, quantity int) error {
	itemIdentifier := ItemIdentifier{
		category: category,
		brand:    brand,
	}
	item := p.items[itemIdentifier]

	if item == nil {
		return errors.New("Item not exist")
	}

	item.quantity += quantity
	fmt.Println("Item inventory added successfully")
	return nil
}

func (p *Platform) AddItemToCart(username string, brand Brand, category Category, quantity int) error {
	itemIdentifier := ItemIdentifier{
		category: category,
		brand:    brand,
	}
	user := p.users[username]
	item := p.items[itemIdentifier]

	if user == nil {
		return errors.New("User not exist")
	}

	if item == nil {
		return errors.New("Item not exist")
	}

	if item.quantity < quantity {
		return fmt.Errorf("out of stock for quantity %d, has only %d items", quantity, item.quantity)
	}

	cartItem := &CartItem{
		item:     item,
		quantity: quantity,
	}

	user.cart = append(user.cart, cartItem)
	fmt.Println("Item added to cart successfully")
	return nil
}

func (p *Platform) RemoveItemFromCart(username string, brand Brand, category Category, quantity int) error {
	user := p.users[username]
	if user == nil {
		return errors.New("User not exist")
	}

	for i, cartItem := range user.cart {
		if cartItem.item.brand == brand && cartItem.item.category == category {
			if cartItem.quantity < quantity {
				return errors.New("Insufficient quantity in cart")
			}

			if cartItem.quantity == quantity {
				user.cart = append(user.cart[:i], user.cart[i+1:]...)
			} else {
				cartItem.quantity -= quantity
			}
			fmt.Println("Item removed from cart successfully")
			return nil
		}
	}
	return errors.New("Item not found in cart")
}

func (p *Platform) Checkout(username string) error {
	user := p.users[username]
	if user == nil {
		return errors.New("User not exist")
	}

	totalAmount := 0

	// validating if out of stock or insufficient balance

	if len(user.cart) == 0 {
		return errors.New("No items in cart")
	}

	for _, cartItem := range user.cart {
		item := cartItem.item
		if item.quantity < cartItem.quantity {
			return errors.New("out of stock")
		}

		totalAmount += item.price * cartItem.quantity
	}

	if user.balance < totalAmount {
		return errors.New("insufficient balance")
	}

	for _, cartItem := range user.cart {
		item := cartItem.item
		item.quantity -= cartItem.quantity
	}

	user.balance = user.balance - totalAmount

	fmt.Println("User purchase completed successfully")
	fmt.Printf("New wallet balance: %d\n", user.balance)
	return nil
}

func (p *Platform) GetUserCart(username string) error {
	user := p.users[username]
	if user == nil {
		return errors.New("User not exist")
	}

	fmt.Printf("%v Cart\n", user.username)

	for _, cartItem := range user.cart {
		item := cartItem.item
		fmt.Printf("%v item of quantity - %v, price - %v, total - %v\n", cartItem, cartItem.quantity, item.price, cartItem.quantity*item.price)
	}
	return nil
}

func (p *Platform) SearchItems(brands []Brand, categories []Category) error {
	items := make([]*Item, 1)
	for _, brand := range brands {
		for _, category := range categories {
			itemIdentifier := ItemIdentifier{
				brand:    brand,
				category: category,
			}
			if item, ok := p.items[itemIdentifier]; ok {
				items = append(items, item)
			}
		}
	}

	for _, item := range items {
		if item != nil {
			fmt.Printf("%v brand - %v category - %v price - %v remaining quantity\n", item.brand, item.category, item.price, item.quantity)
		}
	}
	return nil
}
