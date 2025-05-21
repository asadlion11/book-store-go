package models


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