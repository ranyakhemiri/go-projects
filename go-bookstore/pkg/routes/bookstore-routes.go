package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	// create a new book
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
	// get the list of books
	router.HandleFunc("/book/", controllers.getBook).Methods("GET")
	// get a book by id
	router.HandleFunc("/book/{id}", controllers.getBookById).Methods("GET")
	// update a book by id
	router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT")
	// delete a book by id
	router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE")
}
