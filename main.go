package main 


import (

	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"		
)

type Movie struct {
	ID string json: "id"
	Isbn string  json "Isbn"
	Title string json: "title"
	Dorector *Director json: "director"

}
type Director struct{
	firstname string json: "firstname"
	lastname string json: "lastname"

}

var movies[]Movie


func main() {
	r := mux.newRouter()

	movies = append(movies, Movie(ID:"1", Isbn:"456678", Title: "The HorseMen", Director: &Director(firstname: "John", Lastname : "Doe")))
	movies = append(movies, Movie(ID:"2", Isbn:"456554", Title: "The GO Boys", Director: &Director(firstname: "Ademola", Lastname : "Adebowale")))

	//Create all the routes with mux router
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	// create Function to start the server
	 fmt.Printf("Starting server ar port 8080")
	 log.Fatal(http.ListenAndServe(":8000", r))

}