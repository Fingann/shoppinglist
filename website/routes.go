package website

import (
	"fmt"
	"net/http"
)

//HomePage is the main site of the application
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
