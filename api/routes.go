package api

import (
	"encoding/json"
	"io"
	"net/http"
	"shoppinglist/database"
	"shoppinglist/models"

	"github.com/gorilla/mux"
)

var (
	repo    database.Repository        = database.NewItemListDB()
	elastic database.ElasticRepository = database.NewElasticDB("database-shoppinglist")
)

type itemList []models.Item

// All fetches every record from the db
func All(w http.ResponseWriter, r *http.Request) {

	itemList, err := repo.FindAll()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(itemList)
}

//Single fetches a single document based on id
func Single(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	item, err := repo.Find(id)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}

//Index a document
func Index(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var req models.ItemDTO
	err := dec.Decode(&req)
	elastic.Index(req)

	item, err := repo.Find(r.URL.Query().Get("id"))
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}

//Search for an item in the database
func Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	item, err := elastic.Search(query)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}
