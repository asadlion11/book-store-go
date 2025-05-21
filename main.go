package main

import (
	"encoding/json"
	"log" // simple logging of errors and info
	"math/rand"
	"net/http" // HTTP server and client implementations
	"slices"
	"strconv"

	"github.com/gorilla/mux" // third‑party router for advanced URL matching
)

//Book Model
//Book struct
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Aothor struct
type Author struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

//Init books var as a slice book struct
var books []Book //Book is the slice type and its struct

// Main entry point
func main() {
	//books data //Mock Data
	//book 1
	book1 := Book{
		ID: "1",
		Title: "Yaa Qaatay Burcadkeygii",
		Author: &Author{
			FirstName: "Shaafi Abdi",
			LastName: "Abdi",
		},
	}
	//Append to the books
	books = append(books, book1)
	//book 2
	book2 := Book{
		ID: "2",
		Title: "Dabin",
		Author: &Author{
			FirstName: "Ahmed",
			LastName: "Mohamed",
		},
	}
	//Append to the books
	books = append(books, book2)
	//book 3
	book3 := Book{
		ID: "3",
		Title: "Ninka Ugu Taajirsan Baabil",
		Author: &Author{
			FirstName: "Abyan",
			LastName: "Ahmed",
		},
	}

	//Append to the books
	books = append(books, book3)

	//Init Router
	r := mux.NewRouter()
	//Route Handlers //EndPoints //APIs
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", addBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	//Run Server //Fatal(from log package): if fails
	log.Fatal(http.ListenAndServe(":8000", r))
}

//Functions //Handlers
//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json") 
	 // Encode and write the 'books' data as JSON //returning data as JSON
	json.NewEncoder(w).Encode(books) //books(slice) is the exactlt returning data
}

//Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //get params(params is the parameters of the request/api)
	//Loop throgh books & find with id
	for _, book := range books{
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	} 
	// json.NewEncoder(w).Encode(&Book{})
	http.Error(w, "Book not found", http.StatusNotFound)
	
}

//Create new book //add new book
func addBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	var book Book
	// Read the data sent from the user and put it into the book
	_ = json.NewDecoder(r.Body).Decode(&book) // Take JSON from the request(body) and initialize the book
	book.ID = strconv.Itoa(rand.Intn(1000000)) //not save in production
	books = append(books, book)
	//Returning data is the new book
	json.NewEncoder(w).Encode(book)

}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	//deleting the previos and adding new one
	// w.Header().Set("Content-type", "application/json")
	// params := mux.Vars(r)
	// for i, book := range books {
	// 	if book.ID == params["id"] {
	// 		books = slices.Delete(books, i, i+1)
	// 		var book Book
	// 		_ = json.NewDecoder(r.Body).Decode(&book)
	// 		books = append(books, book)
	// 		json.NewEncoder(w).Encode(book)
	// 		return
	// 	}
	// }

	//Update only title and author keeo the ID
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			var updatingFields Book
			_ = json.NewDecoder(r.Body).Decode(&updatingFields)
			// Only update the title and author, keep the ID the same
			books[i].Title = updatingFields.Title
			books[i].Author = updatingFields.Author
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
		// If no matching book is found, return an empty object or 404
		http.NotFound(w, r)

}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Find the index of the book to delete
	for i, book := range books {
		if book.ID == params["id"] {
			// Remove the book using slices.Delete (Go 1.21+)
			books = slices.Delete(books, i, i+1)
			break
		}
	}
	// If no matching book is found, return an empty object or 404
	http.Error(w, "Book not found", http.StatusNotFound)
}



