package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"burakforum/services"

	"github.com/gorilla/mux"
)

func GetPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	check, err := services.GetpostID(postID)
	if err != nil {
		log.Println("CONTROLLERS PAKETİNDE HATA VAR. GETPOSTID")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if check == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(check)
}
