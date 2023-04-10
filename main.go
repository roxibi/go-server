package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Print(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST req successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}
// HANDLER HAS A RESPONSE AND A REQUEST AS ARGUMENTS
// IF THE PATH OR THE METHOD DEFINED FOR THIS HANDLER ARE INCORRECT, WE RETURN ERRORS FOR THAT; OTHERWISE WE JUST EXECUTE THE HANDLER
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //TELL GO WHERE TO FIND INDEX HTML
	http.Handle("/", fileServer)                        // ROOT ROUTE-NOT A SIMPLE FUNC
	http.HandleFunc("/form", formHandler)               // OTHER ROUTES
	http.HandleFunc("/hello", helloHandler)             // PARAMS=ROUTE AND HANDLER FUNCTION

	fmt.Printf("Server started at port 8080 \n") //TEST SERVER

	//LISTEN AND SERVE DOESNT RETURN ANYTHING UNLESS IT'S AN ERROR - IT JUST STARTS TO RUN IN THE BACKGROUND; SO WE SAVE THE ERROR IN err AND PRINT IT IF IT'S NOT NIL
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
