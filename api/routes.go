package api

import (
	"encoding/json"
	"net/http"
	"shoppinglist/database"
)

// All fetches every record from the db
func All(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.Items)
}

//Single fetches a single document based on id
func Single(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.Items)
}
