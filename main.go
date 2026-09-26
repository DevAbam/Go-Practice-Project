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
			services.ViewProducts()

		case 3:
			var searchproductID int
			fmt.Println("Enter the id of the product to search")
			fmt.Scanln(&searchproductID)
			prod, err := services.SearchProduct(searchproductID)
			if err != nil {
				fmt.Printf("product with id %d not found", searchproductID)
				continue
			}
			fmt.Println("product found: ", prod)

		case 4:
			services.UpdateStock()

		case 5:
			services.RemoveProduct()

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid menu choice.")
		}
	}
}