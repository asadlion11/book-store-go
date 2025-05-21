package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
	"my-first-server/models"
	"my-first-server/utils"
)

//Functions //Handlers
//Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json") 
	 // Encode and write the 'books' data as JSON //returning data as JSON
	json.NewEncoder(w).Encode(utils.Books) //books(slice) is the exactlt returning data
}

//Get book
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //get params(params is the parameters of the request/api)
	//Loop throgh books & find with id
	for _, book := range utils.Books{
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	} 
	json.NewEncoder(w).Encode(&models.Book{})

}

//Create new book //add new book
func AddBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	// Read the data sent from the user and put it into the book
	_ = json.NewDecoder(r.Body).Decode(&book) // Take JSON from the request(body) and initialize the book
	book.ID = strconv.Itoa(rand.Intn(1000000)) //not save in production
	utils.Books = append(utils.Books, book)
	// //Returning data is the new book
	// json.NewEncoder(w).Encode(book)
	json.NewEncoder(w).Encode(utils.Books) 

}

//Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Update only title and author keeo the ID
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for i, book := range utils.Books {
		if book.ID == params["id"] {
			var updatingFields models.Book
			_ = json.NewDecoder(r.Body).Decode(&updatingFields)
			// Only update the title and author, keep the ID the same
			utils.Books[i].Title = updatingFields.Title
			utils.Books[i].Author = updatingFields.Author
			json.NewEncoder(w).Encode(utils.Books[i])
			return
		}
	}
}

//Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Tell the client we're returning JSON
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Find the index of the book to delete
	for i, book := range utils.Books {
		if book.ID == params["id"] {
			// Remove the book using slices.Delete (Go 1.21+)
			utils.Books = slices.Delete(utils.Books, i, i+1)
			break
		}
	}
	json.NewEncoder(w).Encode(utils.Books)
}