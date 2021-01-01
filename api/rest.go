package api

import (
	"encoding/json"
	"net/http"
	"shoppinglist/models"
	"strconv"
	"strings"
	"time"
)

// Rest API type
type Rest struct{}

// Items simulate rows fetched from a database
var items []models.Item

func init() {
	items = []models.Item{
		models.Item{Name: "Paprika", Desc: "skal ha fin fasong", Quantity: 2, Added: time.Now(), Completed: time.Now().Add(time.Hour * 4)},
		models.Item{Name: "is", Desc: "ben & jerries", Quantity: 1, Added: time.Now()},
	}
}

// All fetches every record from the db
func (Rest) All(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

//Single fetches a single document based on id
func (Rest) Single(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	lastIndex := len(paths) - 1
	index, err := strconv.Atoi(paths[lastIndex])
	if err != nil {
		http.Error(w, "Unable to parse index from path", http.StatusBadRequest)
		return
	}
	if index >= 0 && index < len(items) {
		json.NewEncoder(w).Encode(items[index])

	} else {
		http.Error(w, "Could not find the id", http.StatusNotFound)
	}
}

// SetupRequests sets up all rest enpoints for the api
//func SetupRequests() {
//	http.HandleFunc("/api/items", allItems)
//}
