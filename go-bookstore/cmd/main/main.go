package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ranyakhemiri/go-bookstore/pkg/routes"
)

// create the server
// precise where our routes reside
func main() {
	// creating a new router
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	// listen on port 9010, if there is an error, log it
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
