package main

import (
	"fmt"
	"log"
	"net/http"
	"shoppinglist/api"
	"shoppinglist/website"
)

const (
	// Host which the server is running on
	Host string = "127.0.0.1"

	// Port used by the server
	Port string = "10000"
)

func setupAPIRoutes() {
	http.HandleFunc("/api/items", api.All)
}
func setupWebsiteRoutes() {
	http.HandleFunc("/", website.HomePage)
}

func main() {
	setupAPIRoutes()
	setupWebsiteRoutes()

	log.Println("Strating the server on port: ", Port)
	address := fmt.Sprint(Host, ":", Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
