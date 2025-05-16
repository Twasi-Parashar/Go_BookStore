package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct for book details
type Book struct {
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Saving File Function
func saveToFile(book Book) {
	//Reading existing file to append
	var books []Book
	fileData, err := os.ReadFile("data.json")
	if err == nil && len(fileData) > 0 {
		err = json.Unmarshal(fileData, &books)
		if err != nil {
			fmt.Println("Error in unmarshling existing JSON:", err)
			return
		}
	}

	//Appending new book data
	books = append(books, book)

	//Updating the file
	bookJson, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		fmt.Println("Error in marshling book data", err)
		return
	}

	//writing into the file
	err = os.WriteFile("data.json", bookJson, 0644)
	if err != nil {
		fmt.Println("Error in writing to file:", err)
	}

	fmt.Println("Book data saved Successfully.")
}

func main() {
	//Menu
	fmt.Println("BOOK STORE SYSTEM")
	fmt.Println("1. View Books")
	fmt.Println("2. Add Book")
	fmt.Println("Enter the choice: ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		os.Exit(1)
	}

	//Switch case
	switch choice {
	case 1:
		viewBooks()
	case 2:
		addBooks()
	default:
		fmt.Println("Error in choice")
	}

}

// ViewBooks function
func viewBooks() {
	//Acessing the file
	fileData, err := os.ReadFile("data.json")
	if err != nil || len(fileData) == 0 {
		fmt.Println("No books found or error in reading file")
		return
	}

	//Unmarshaling
	var books []Book
	err = json.Unmarshal(fileData, &books)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	//Displaying
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}
	fmt.Println("Book List")
	for i, book := range books {
		fmt.Printf("%d. %s by %s - ₹%.2f\n", i+1, book.Name, book.Author, book.Price)
	}
}

// AddBook Function
func addBooks() {
	//Taking input of book details
	fmt.Println("Enter Book details")
	fmt.Println("Enter name of the book: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter the author of the book: ")
	var author string
	fmt.Scan(&author)
	fmt.Println("Enter the price of the book: ")
	var price float64
	fmt.Scan(&price)

	//putting in struct
	book := Book{
		Name:   name,
		Author: author,
		Price:  price,
	}

	//Saving it to json file
	saveToFile(book)
}
