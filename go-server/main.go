package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}	
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	}
	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}
 
func main(){
	// Registering the handler
	// The handler is the default server mux, 
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	// Starting the server
	fmt.Println("Server is listening on port 8080")
	if err:= http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}


}	