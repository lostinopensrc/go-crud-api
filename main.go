package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lostinopensrc/go-crud-api/handlers"
	"github.com/lostinopensrc/go-crud-api/schema"
)

func main() {

	// initialize values for Movies slice
	schema.Movies = append(schema.Movies, schema.Movie{Id: "1", Isbn: "123456", Title: "The phoenix Project", Director: &schema.Director{FirstName: "Dan", LastName: "Lee"}})
	schema.Movies = append(schema.Movies, schema.Movie{Id: "2", Isbn: "789101", Title: "The Unicorn Project", Director: &schema.Director{FirstName: "Dan", LastName: "Lee"}})
	router := mux.NewRouter() // declared router object

	router.HandleFunc("/movies", handlers.Get_Movies).Methods("GET")           // get Movies
	router.HandleFunc("/movies", handlers.Create_Movie).Methods("POST")        // create Movies
	router.HandleFunc("/movies/{id}", handlers.Get_Movie).Methods("GET")       // Get Movie by Id
	router.HandleFunc("/movies/{id}", handlers.Delete_Movie).Methods("DELETE") //  Delete Movie
	router.HandleFunc("/movies/{id}", handlers.Update_Movie).Methods("PUT")    //  Update Movie

	fmt.Println("Starting server at 8000 port")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
