package main

import "fmt"

func main() {

	var book Product = Product{
		ID:               "353543",
		name:             "Name",
		shortDescription: "This is a test product",
		price:            100,
	}

	pen := createProduct("23423", "Sharikiviy", "It is a test", 12)
	outputResult(book)
	outputResult(pen)

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
	price            int
}

// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
func createProduct(ID string, name string, description string, price int) (product Product) {
	product = Product{
		ID:               ID,
		name:             name,
		shortDescription: description,
		price:            price,
	}
	return
}

// 3) Add a "connected function" that outputs the data + call that function from inside "main"
func outputResult(product Product) {
	fmt.Println(product)
}

// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.

// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
