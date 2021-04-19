package handlers

import (
	"log"
	"net/http"
)

func GetBookmarkListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit Create Bookmark")
	w.WriteHeader(http.StatusOK)
}
