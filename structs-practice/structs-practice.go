package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//creating struct
type Product struct {
	id string
	title string
	description string
	price float64
}

//storing user data in a file
func (prod *Product) store() {
	file,_ :=os.Create(prod.id + ".txt")
	content := fmt.Sprintf(
		"ID: %v\n Title: %v\n Description: %v\n Price: $%.2f\n", 
		prod.id,prod.title,prod.description,prod.price)
	file.WriteString(content)
	file.Close()
}

//printing data
func (prod *Product) printData () {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: $%.2f\n", prod.price)
}


//creating product struct
func NewProduct (id string, t string, d string, p float64) *Product {

	return  &Product{id, t, d, p}
}


func main() {
	createdProduct := getProduct()
	createdProduct.printData()
	createdProduct.store()
}


//getting  user inputs 
func getProduct() *Product {
	fmt.Print("Please enter the product\n")
	fmt.Println("--------------------------")
	reader := bufio.NewReader(os.Stdin)
	idInput := readerUserInput(reader, "Product ID: ")
	titleInput:= readerUserInput(reader, "Product Title: ")
	descriptionInput:= readerUserInput(reader, "Product Description: ")
	priceInput:= readerUserInput(reader, "Product Price: ")
	priceValue,_ := strconv.ParseFloat(priceInput, 64)
	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)
	return product
}

//Cleaning user input
func readerUserInput( reader *bufio.Reader, promptText string) string{
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}



// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.