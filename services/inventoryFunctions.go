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

func ViewProduct() {
}

func SearchProduct() {
}

func UpdateStock() {
}

func RemoveProduct() {
}

func InventoryValue() {
}

func LowStock() {
}