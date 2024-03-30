package main

func main() {
	platform := NewPlatform()

	platform.AddUser("user1", "address1", 1000)
	platform.AddUser("user2", "address2", 2000)

	platform.AddItem(Amul, Curd, 100)
	platform.AddItem(Amul, Milk, 50)
	platform.AddItem(Amul, Cheese, 200)

	platform.AddItem(Nestle, Curd, 150)
	platform.AddItem(Nestle, Milk, 70)
	platform.AddItem(Nestle, Cheese, 250)

	platform.AddItemInventory(Amul, Curd, 100)
	platform.AddItemInventory(Amul, Milk, 50)
	platform.AddItemInventory(Amul, Cheese, 200)

	platform.AddItemToCart("user1", Amul, Curd, 2)
	platform.AddItemToCart("user1", Amul, Milk, 3)

	platform.AddItemToCart("user2", Nestle, Curd, 1)
	platform.AddItemToCart("user2", Nestle, Milk, 2)

	platform.GetUserCart("user1")
	platform.RemoveItemFromCart("user1", Amul, Curd, 1)
	platform.GetUserCart("user1")
	platform.GetUserCart("user2")

	platform.Checkout("user1")
	platform.Checkout("user2")

	platform.SearchItems([]Brand{Amul, Nestle}, []Category{Curd, Milk})

}
