package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// net/http package is used to create a web server and handle HTTP requests.
// encoding/json package is used to encode and decode JSON data (because we are using Postman, we want to be able to decode and encode JSON data).
// log package is used to log errors and messages.
// math/rand package is used to generate random numbers (ids for movies)
// strconv package is used to convert strings to other types.
// gorilla/mux is used to create an HTTP router.

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Slice of movies
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we want to encode the response as JSON
	// the response will be the list of movies
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete that movie using append() method
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// return the rest of the movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create a movie object
	var movie Movie
	// decode the request body and store it in the movie object
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generate a random id for the movie
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	// append the movie to the list of movies
	movies = append(movies, movie)
	// return the list of movies
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type of the response
	w.Header().Set("Content-Type", "application/json")
	// get the id from the request parameters
	params := mux.Vars(r)
	// delete the movie with the given id
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie                            // create a new movie object
			_ = json.NewDecoder(r.Body).Decode(&movie) // decode the request body and store it in the movie object
			movie.ID = params["id"]
			// add the new movie to the list of movies
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	// to declare a Movie struct : StructName{ FieldName: value, FieldName: value, ... }
	movies = append(movies, Movie{ID: "1", Isbn: "4324029", Title: "Harry Potter and the Philosopher's Stone", Director: &Director{Firstname: "Chris", Lastname: "Columbus"}})
	movies = append(movies, Movie{ID: "2", Isbn: "24029302", Title: "Harry Potter and the Chamber of Secrets", Director: &Director{Firstname: "Chris", Lastname: "Columbus"}})
	// Creating 5 different routes ...
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")

	// Start the server
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
