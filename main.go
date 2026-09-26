package main

import (
	"fmt"
	"ingabam.com/inventoryapp/services"
)

func main() {
	var menuChoice int

	menus := []string{
		"Add Product",
		"View Product",
		"Search Product",
		"Update Stock",
		"Remove Product",
		"Inventory Value",
		"Low Stock",
		"Exit",
	}

	for {
		fmt.Println("\nInventory Menu: Select an option to continue")

		for i, menu := range menus {
			fmt.Println(i+1, menu)
		}

		fmt.Print("Choose an option: ")
		fmt.Scanln(&menuChoice)

		switch menuChoice {
		case 1:
			var productID, quantity int
			var name, category string
			var price float64

			fmt.Print("Enter product ID: ")
			fmt.Scanln(&productID)

			fmt.Print("Enter quantity: ")
			fmt.Scanln(&quantity)

			fmt.Print("Enter name: ")
			fmt.Scanln(&name)

			fmt.Print("Enter category: ")
			fmt.Scanln(&category)

			fmt.Print("Enter price: ")
			fmt.Scanln(&price)

			err := services.AddProduct(
				productID,
				name,
				price,
				quantity,
				category,
			)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Product added successfully!")

		case 2:
			services.ViewProduct()

		case 3:
			services.SearchProduct()

		case 4:
			services.UpdateStock()

		case 5:
			services.RemoveProduct()

		case 6:
			services.InventoryValue()

		case 7:
			services.LowStock()

		case 8:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid menu choice.")
		}
	}
}