package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

// AddProduct function
func AddProduct(id int, name string, price interface{}, stock int) error {
	convertedPrice, ok := price.(float64)
	if !ok {
		return errors.New("price must be a float64 value")
	}

	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: convertedPrice,
		Stock: stock,
	}
	inventory = append(inventory, newProduct)
	return nil
}

// UpdateStock stock
func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProduct function
func SearchProduct(identifier interface{}) (Product, error) {
	switch v := identifier.(type) {
	case int: // Search by id
		for _, product := range inventory {
			if product.ID == v {
				return product, nil
			}
		}
      
	case string: // Search by name
		for _, product := range inventory {
			if product.Name == v {
				return product, nil
			}
		}
      
	default:
		return Product{}, errors.New("identifier must be an int or string")
	}

	return Product{}, errors.New("product not found")
}

// DisplayInventory function
func DisplayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-5s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("----------------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-5d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventory function
func SortInventory(sortBy string) error {
	switch sortBy {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		return errors.New("invalid sort option; use 'price' or 'stock'")
	}
	return nil
}

func main() {
	_ = AddProduct(1, "Laptop", 57654.95, 10)
	_ = AddProduct(2, "Mouse", 465.50, 50)
	_ = AddProduct(3, "Keyboard", 932.75, 30)

	fmt.Println("Initial Inventory:")
	DisplayInventory()

  
	fmt.Println("\nUpdating stock of product with ID 1:")
	_ = UpdateStock(1, 15)
	DisplayInventory()

  
	fmt.Println("\nSearching for product with ID 2:")
	if product, err := SearchProduct(2); err == nil {
		fmt.Printf("Found: %+v\n", product)
	} else {
		fmt.Println(err)
	}
  

	fmt.Println("\nSorting inventory by price:")
	_ = SortInventory("price")
	DisplayInventory()


	fmt.Println("\nSorting inventory by stock:")
	_ = SortInventory("stock")
	DisplayInventory()
}
