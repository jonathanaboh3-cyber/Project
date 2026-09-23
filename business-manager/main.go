package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Customer struct {
	Name  string
	Phone string
}

var products []Product
var sales []Sale
var customers []Customer

func saveProducts() {
	file, err := os.Create("products.txt")

	if err != nil {
		fmt.Println("Error saving products:", err)
		return
	}

	defer file.Close()

	for _, product := range products {
		data := fmt.Sprintf(
			"%s,%.2f,%d\n",
			product.Name,
			product.Price,
			product.Quantity,
		)

		file.WriteString(data)
	}
}

func saveSale(sale Sale) {
	data := fmt.Sprintf(
		"%s,%.2f,%d,%.2f\n",
		sale.ProductName,
		sale.Price,
		sale.Quantity,
		sale.Total,
	)

	file, err := os.OpenFile(
		"sales.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Error saving sale:", err)
		return
	}

	file.WriteString(data)
	file.Close()
}

func loadProducts() {
	data, err := os.ReadFile("products.txt")

	if err != nil {
		fmt.Println("Error reading products:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ",")

		if len(parts) != 3 {
			fmt.Println("Invalid product data:", line)
			continue
		}

		price, err := strconv.ParseFloat(parts[1], 64)

		if err != nil {
			fmt.Println("Invalid price:", parts[1])
			continue
		}

		quantity, err := strconv.Atoi(parts[2])

		if err != nil {
			fmt.Println("Invalid quantity:", parts[2])
			continue
		}

		product := Product{
			Name:     parts[0],
			Price:    price,
			Quantity: quantity,
		}

		products = append(products, product)
	}
}

func loadSales() {
	data, err := os.ReadFile("sales.txt")

	if err != nil {
		fmt.Println("Error reading sales:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ",")

		if len(parts) != 4 {
			fmt.Println("Invalid sale data:", line)
			continue
		}

		price, err := strconv.ParseFloat(parts[1], 64)

		if err != nil {
			fmt.Println("Invalid price:", parts[1])
			continue
		}

		quantity, err := strconv.Atoi(parts[2])

		if err != nil {
			fmt.Println("Invalid quantity:", parts[2])
			continue
		}

		total, err := strconv.ParseFloat(parts[3], 64)

		if err != nil {
			fmt.Println("Invalid total:", parts[3])
			continue
		}

		sale := Sale{
			ProductName: parts[0],
			Price:       price,
			Quantity:    quantity,
			Total:       total,
		}

		sales = append(sales, sale)
	}
}

func addProduct() {
	var productName string

	fmt.Print("Enter product name: ")
	fmt.Scanln(&productName)

	productName = strings.ToLower(productName)

	exists := false

	for _, product := range products {
		if product.Name == productName {
			exists = true
			break
		}
	}

	if exists {
		fmt.Println("Product already exists.")
		return
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
}

func viewProducts() {
	fmt.Println()
	fmt.Println("===== PRODUCTS =====")

	if len(products) == 0 {
		fmt.Println("No products available.")
		return
	}

	for i, product := range products {
		fmt.Println("Product", i+1)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Quantity:", product.Quantity)
		fmt.Println()
	}
}

func processPurchase(index int) {
	product := products[index]

	var quantity int

	for {
		fmt.Print("Enter quantity to buy: ")

		if _, err := fmt.Scanln(&quantity); err != nil {
			fmt.Println("Invalid quantity")
		} else if quantity <= 0 {
			fmt.Println("Quantity must be greater than 0")
		} else {
			break
		}
	}

	if quantity > product.Quantity {
		fmt.Println("Not enough quantity in stock.")
		return
	}

	products[index].Quantity -= quantity

	total := product.Price * float64(quantity)

	sale := Sale{
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
		Total:       total,
	}

	sales = append(sales, sale)

	saveProducts()
	saveSale(sale)

	fmt.Println()
	fmt.Println("===== PURCHASE SUCCESSFUL =====")
	fmt.Println("Product:", product.Name)
	fmt.Println("Quantity:", quantity)
	fmt.Println("Total amount:", total)
}

func searchProduct() {
	var searchName string

	fmt.Print("Enter product name: ")
	fmt.Scanln(&searchName)

	searchName = strings.ToLower(searchName)

	found := false

	for i, product := range products {
		if product.Name == searchName {
			found = true

			fmt.Println()
			fmt.Println("===== PRODUCT FOUND =====")
			fmt.Println("Name:", product.Name)
			fmt.Println("Price:", product.Price)
			fmt.Println("Quantity:", product.Quantity)

			fmt.Println()
			fmt.Println("1. Buy Product")
			fmt.Println("2. Back")

			var choice int

			fmt.Print("Choose an option: ")
			fmt.Scanln(&choice)

			switch choice {

			case 1:
				processPurchase(i)

			case 2:
				return

			default:
				fmt.Println("Invalid option.")
			}

			break
		}
	}

	if !found {
		fmt.Println("Product not found.")
	}
}

func recordSale() {
	var saleProduct string

	fmt.Print("Enter product name: ")
	fmt.Scanln(&saleProduct)

	saleProduct = strings.ToLower(saleProduct)

	found := false

	for i, product := range products {

		if product.Name == saleProduct {
			found = true

			fmt.Println("Product found:", product)

			processPurchase(i)

			break
		}
	}

	if !found {
		fmt.Println("Product not found.")
	}
}

func viewSales() {
	fmt.Println()
	fmt.Println("===== SALES =====")

	if len(sales) == 0 {
		fmt.Println("No sales recorded.")
		return
	}

	for i, sale := range sales {
		fmt.Println("Sale", i+1)
		fmt.Println("Product:", sale.ProductName)
		fmt.Println("Price:", sale.Price)
		fmt.Println("Quantity:", sale.Quantity)
		fmt.Println("Total:", sale.Total)
		fmt.Println()
	}
}

func addCustomer() {
	var customerName string
	var customerPhone string

	fmt.Print("Enter customer name: ")
	fmt.Scanln(&customerName)

	fmt.Print("Enter customer phone: ")
	fmt.Scanln(&customerPhone)

	customer := Customer{
		Name:  customerName,
		Phone: customerPhone,
	}

	customers = append(customers, customer)

	fmt.Println("Customer added successfully.")
}

func viewCustomers() {
	fmt.Println()
	fmt.Println("===== CUSTOMERS =====")

	if len(customers) == 0 {
		fmt.Println("No customers available.")
		return
	}

	for i, customer := range customers {
		fmt.Println("Customer", i+1)
		fmt.Println("Name:", customer.Name)
		fmt.Println("Phone:", customer.Phone)
		fmt.Println()
	}
}

func main() {

	loadProducts()
	loadSales()

	for {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("            BIZFLOW")
		fmt.Println("================================")

		fmt.Println()
		fmt.Println("1. Add Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Search Product")
		fmt.Println("4. Record Sale")
		fmt.Println("5. View Sales")
		fmt.Println("6. add Customer")
		fmt.Println("7. view Customers")
		fmt.Println("8. Goodbye!")
		fmt.Println()

		var choice int

		fmt.Print("Choose an option: ")

		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid option. Please enter a number.")
			continue
		}

		switch choice {

		case 1:
			addProduct()

		case 2:
			viewProducts()

		case 3:
			searchProduct()

		case 4:
			recordSale()

		case 5:
			viewSales()

		case 6:
			addCustomer()

		case 7:
			viewCustomers()

		case 8:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
