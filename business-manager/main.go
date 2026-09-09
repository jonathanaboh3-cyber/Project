package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	fmt.Println("================================")
	fmt.Println("       BUSINESS MANAGER")
	fmt.Println("================================")

	var choice int

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

		fmt.Print("Enter product price: ")
		fmt.Scanln(&productPrice)

		fmt.Println("Product price:", productPrice)

		var productQuantity int

		fmt.Print("Enter product quantity: ")
		fmt.Scanln(&productQuantity)

		fmt.Println("Product quantity:", productQuantity)

		product := Product{
		Name:     productName,
		Price:    productPrice,
		Quantity: productQuantity,
	}

		fmt.Println("Product added:", product)

	case 2:
		fmt.Println("View Products selected")

	case 3:
		fmt.Println("Record Sale selected")

	case 4:
		fmt.Println("View Sales selected")

	case 5:
		fmt.Println("Goodbye!")

	default:
		fmt.Println("Invalid option")
	}
}