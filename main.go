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
		fmt.Println("---------------------------------------------------------")
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
			var productID int
			var amount int
			var choice int

			fmt.Println("Enter product ID:")
			fmt.Scanln(&productID)

			fmt.Println("1. Add Stock")
			fmt.Println("2. Remove Stock")
			fmt.Scanln(&choice)

			fmt.Println("Enter amount:")
			fmt.Scanln(&amount)

			var operation string

			switch choice {
			case 1:
				operation = "add"

			case 2:
				operation = "remove"

			default:
				fmt.Println("Invalid choice")
				continue
			}

			err := services.UpdateStock(productID, amount, operation)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Stock updated successfully")

		case 5:
			var removeproductID int
			fmt.Println("Enter the id of the product to remove")
			fmt.Scanln(&removeproductID)
			err := services.RemoveProduct(removeproductID)
			if err!= nil{
				fmt.Printf("product with id %d not deleted", removeproductID)
				continue
			}
			fmt.Println("Product deleted succesfully")

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid menu choice.")
		}
	}
}