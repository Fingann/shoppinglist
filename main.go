package main

import (
	"fmt"
	"log"
	"net/http"
	"shoppinglist/api"
)

const (
	// Host which the server is running on
	Host string = "127.0.0.1"

	// Port used by the server
	Port string = "10000"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// ItemsAPI represents functions to fetch items
type itemsHandler interface {
	All(w http.ResponseWriter, r *http.Request)
	Single(w http.ResponseWriter, r *http.Request)
}

func handleRequests(handler itemsHandler) {

	http.HandleFunc("/api/items", itemsHandler.All)

	log.Println("Strating the server on port: ", Port)
	address := fmt.Sprint(Host, ":", Port)
	log.Fatal(http.ListenAndServe(address, nil))
}

func main() {
	var handler itemsHandler
	handler = api.Rest{}

	handleRequests(handler)
}
