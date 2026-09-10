package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

var products []Product

func main() {
	fmt.Println("================================")
	fmt.Println("       BUSINESS MANAGER")
	fmt.Println("================================")

	var choice int

	for {
		fmt.Println()
		fmt.Println("1. Add Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Record Sale")
		fmt.Println("4. View Sales")
		fmt.Println("5. Exit")
		fmt.Println()

		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var productName string

			fmt.Print("Enter product name: ")
			fmt.Scanln(&productName)

			fmt.Println("Product name:", productName)

			var productPrice float64

			for {
				fmt.Print("Enter product price: ")

				if _, err := fmt.Scanln(&productPrice); err != nil {
					fmt.Println("Invalid price")
				} else {
					break
				}
			}

			fmt.Println("Product price:", productPrice)

			var productQuantity int

			for {
				fmt.Print("Enter product quantity: ")

				if _, err := fmt.Scanln(&productQuantity); err != nil {
					fmt.Println("Invalid quantity")
				} else if productQuantity <= 0 {
					fmt.Println("Quantity must be greater than 0")
				} else {
					break
				}
			}

			fmt.Println("Product quantity:", productQuantity)

			product := Product{
				Name:     productName,
				Price:    productPrice,
				Quantity: productQuantity,
			}

			products = append(products, product)

			fmt.Println("Product added:", product)

		case 2:
			fmt.Println()
			fmt.Println("===== PRODUCTS =====")

			if len(products) == 0 {
		        fmt.Println("No products available.")
            }
			
			for i, product := range products {
				fmt.Println("Product", i+1)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println()
			}
		case 3:
			fmt.Println("Record Sale selected")

		case 4:
			fmt.Println("View Sales selected")

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}