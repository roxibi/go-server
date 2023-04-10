package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //use * to access the value at the address we receive. POSSIBLE REASON: BETTER TO USE A POINTER AS A BASIC VALUE THAN A STRUCT
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Print(w, "ParseForm() err:%v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST req successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "Name =%s \n", name)
// 	fmt.Fprintf(w, "Address = %s \n", address)
// }
// // HANDLER HAS A RESPONSE AND A REQUEST AS ARGUMENTS
// // IF THE PATH OR THE METHOD DEFINED FOR THIS HANDLER ARE INCORRECT, WE RETURN ERRORS FOR THAT; OTHERWISE WE JUST EXECUTE THE HANDLER
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "method is not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "hello")
// }

func main() {

	r := mux.NewRouter()


	movies = append(movies, Movie{ID:"1", Isbn: "438227", Title: "Movie One", Director : &Director{Firstname:"Alice", Lastname: "of Wonderland"}}) //USE "&"" TO SEND ONLY THE ADDRESS OF THE VALUE; IT WILL LATER BE USED TO GET TO THE ACTUAL VALUE WITH *

	movies = append(movies, Movie{ID:"2", Isbn: "438116", Title: "Movie Two", Director : &Director{Firstname:"Carla", Lastname: "Bruni"}})
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fileServer := http.FileServer(http.Dir("./static")) //TELL GO WHERE TO FIND INDEX HTML
	http.Handle("/", fileServer)           // ROOT 

	fmt.Printf("Server started at port 8080 \n") //TEST SERVER

	
	log.Fatal(http.ListenAndServe(":8080", r))
	

}
