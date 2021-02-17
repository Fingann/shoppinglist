package main

import (
	"log"
	"net/http"
	"shoppinglist/api"
	"shoppinglist/website"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
)

const (
	// Address the host and port which the server is running on
	Address string = "127.0.0.1:10000"
)

func setupAPIRoutes(r *mux.Router) {
	r.Path("/items").HandlerFunc(api.All)
	r.Path("/item/{id:[0-9]+}").HandlerFunc(api.Single)
	r.Path("/item").HandlerFunc(api.Index).Methods("POST")
	r.Path("/item/search/{query:[\\w]+").HandlerFunc(api.Search)
}

func setupWebsiteRoutes(r *mux.Router) {
	r.Path("/").HandlerFunc(website.HomePage)
}

func main() {
	es, _ := elasticsearch.NewDefaultClient()

	//log.Println(es.Version())
	log.Println(es.Info())
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	websiteRouter := r.NewRoute().Subrouter()

	setupAPIRoutes(apiRouter)
	setupWebsiteRoutes(websiteRouter)

	srver := &http.Server{
		Handler: r,
		Addr:    Address,
	}
	log.Println("Strating the server on port: ", srver.Addr)

	log.Fatal(srver.ListenAndServe())
}
