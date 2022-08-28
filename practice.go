package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	createdProduct := getProduct()

	createdProduct.outputResult()

}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
type Product struct {
	ID               string
	name             string
	shortDescription string
	price            float64
}

func (product *Product) outputResult() {
	fmt.Printf("ID: %v\n Name: %v\n Description: %v\n Price: $%.2f\n", product.ID, product.name, product.shortDescription, product.price)
}

func createProduct(ID string, name string, description string, price float64) (product *Product) {
	product = &Product{
		ID:               ID,
		name:             name,
		shortDescription: description,
		price:            price,
	}
	return
}

func getProduct() *Product {
	fmt.Println("PLease enter the product data")
	fmt.Println("--------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	nameInput := readUserInput(reader, "Product name: ")
	descrition := readUserInput(reader, "Descrition: ")
	priceInput := readUserInput(reader, "Product price")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := createProduct(idInput, nameInput, descrition, priceValue)
	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Println(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
