package router

import (

	"my-first-server/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router{
	//Init Router
	r := mux.NewRouter()
	//Route Handlers //EndPoints //APIs
	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")
	return r
}