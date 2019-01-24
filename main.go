package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}



func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	fmt.Println("Listening at :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
