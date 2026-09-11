package main

import (
	"fmt"
	"os"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Sale struct {
	ProductName string
	Price       float64
	Quantity    int
	Total       float64
}

var products []Product
var sales []Sale

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

			exists := false

			for _, product := range products {
				if product.Name == productName {
					exists = true
					break
				}
			}

			if exists {
				fmt.Println("Product already exists.")
				continue
			}

			var productPrice float64

			for {
				fmt.Print("Enter product price: ")

				if _, err := fmt.Scanln(&productPrice); err != nil {
					fmt.Println("Invalid price")
				} else if productPrice <= 0 {
					fmt.Println("Price must be greater than 0")
				} else {
					break
				}
			}

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

			product := Product{
				Name:     productName,
				Price:    productPrice,
				Quantity: productQuantity,
			}

			products = append(products, product)

			data := fmt.Sprintf(
				"%s,%.2f,%d\n",
				product.Name,
				product.Price,
				product.Quantity,
			)

			file, err := os.OpenFile(
				"products.txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY,
				0644,
			)

			if err != nil {
				fmt.Println("Error saving product:", err)
			} else {
				file.WriteString(data)
				file.Close()
				fmt.Println("Product saved to file.")
			}

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
			var saleProduct string

			fmt.Print("Enter product name: ")
			fmt.Scanln(&saleProduct)

			found := false

			for i, product := range products {
				if product.Name == saleProduct {
					found = true

					fmt.Println("Product found:", product)

					var saleQuantity int

					for {
						fmt.Print("Enter quantity to buy: ")

						if _, err := fmt.Scanln(&saleQuantity); err != nil {
							fmt.Println("Invalid quantity")
						} else if saleQuantity <= 0 {
							fmt.Println("Quantity must be greater than 0")
						} else {
							break
						}
					}

					if saleQuantity > product.Quantity {
						fmt.Println("Not enough quantity in stock.")
					} else {
						products[i].Quantity =
							products[i].Quantity - saleQuantity

						totalAmount :=
							product.Price * float64(saleQuantity)

						fmt.Println("Quantity bought:", saleQuantity)
						fmt.Println("Total amount:", totalAmount)

						sale := Sale{
							ProductName: product.Name,
							Price:       product.Price,
							Quantity:    saleQuantity,
							Total:       totalAmount,
						}

						sales = append(sales, sale)

						fmt.Println("Sale recorded successfully.")
					}

					break
				}
			}

			if !found {
				fmt.Println("Product not found.")
			}

		case 4:
			fmt.Println()
			fmt.Println("===== SALES =====")

			if len(sales) == 0 {
				fmt.Println("No sales recorded.")
			}

			for i, sale := range sales {
				fmt.Println("Sale", i+1)
				fmt.Println("Product:", sale.ProductName)
				fmt.Println("Price:", sale.Price)
				fmt.Println("Quantity:", sale.Quantity)
				fmt.Println("Total:", sale.Total)
				fmt.Println()
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}