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
	createdProduct.store()
}

type Product struct {
	ID               string
	name             string
	shortDescription string
	price            float64
}

func (product *Product) store() {
	file, _ := os.Create(product.name + product.ID + ".txt")

	content := fmt.Sprintf("ID: %v\nName: %v\nDescription: %v\nPrice: %v\n",
		product.ID,
		product.name,
		product.shortDescription,
		product.price)

	file.WriteString(content)
	file.Close()
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
