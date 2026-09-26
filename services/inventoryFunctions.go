package services

import (
	"fmt"

	"ingabam.com/inventoryapp/entity"
	"ingabam.com/inventoryapp/repository"
)

func AddProduct(
	productID int,
	name string,
	price float64,
	quantity int,
	category string,
) error {

	if productID <= 0 {
		return fmt.Errorf("product ID must be greater than 0")
	}

	if price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	if quantity < 0 {
		return fmt.Errorf("quantity cannot be negative")
	}

	if name == "" {
		return fmt.Errorf("product name cannot be empty")
	}

	if category == "" {
		return fmt.Errorf("product category cannot be empty")
	}

	for _, prod := range repository.Products {
		if prod.ID == productID {
			return fmt.Errorf("product ID already exists")
		}
	}

	newProduct := entity.Product{
		ID:       productID,
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}

	repository.Products = append(repository.Products, newProduct)

	return nil
}

func ViewProducts() {
	tableheaders := []string{"ID", "NAME", "PRICE", "STOCK", "CATEGORY"}
	for _ , header := range tableheaders{
		fmt.Printf("%-10s  ", header)
	}
	fmt.Println("")
	fmt.Println("--------------------------------------------------------------------")

	if len(repository.Products) == 0{
		fmt.Println("No products found")
	}

	for _ , prod := range repository.Products{
		fmt.Printf("%-10d  %-10s  %-10.2f  %-10d %-10s\n", prod.ID, prod.Name, prod.Price, prod.Quantity, prod.Category )
	}
}

func SearchProduct(productId int) (entity.Product, error) {
	for _ , prod := range repository.Products{
		if prod.ID == productId{
			return prod, nil
		}
	}
	
	return entity.Product{} , fmt.Errorf("Product not found")
}

func UpdateStock() {
}

func RemoveProduct() {
}

