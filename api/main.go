package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/references/{id}", ReferencesHandler)
	http.Handle("/", r)
	fmt.Print("Something")
}

func ReferencesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}
