package main

import (
	"Brice-Ridings/leaftea_bookmarks_api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("Bookmark Main: ")
	log.SetFlags(0)

	r := mux.NewRouter()
	r.HandleFunc("/bookmarks", handlers.CreateBookmarkHandler).Methods("POST")
	r.HandleFunc("/bookmarks", handlers.GetBookmarkListHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
