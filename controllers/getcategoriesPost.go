package controllers

import (
	"forumbackend/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategoriesPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SelectPostLang, exists := vars["id"]
	if !exists {
		log.Printf("ID not found in URL") // Debug log
		http.Error(w, "Missing ID in URL", http.StatusBadRequest)
		return
	}

	log.Printf("ID received: %s", SelectPostLang) // Debug log

	post, err := services.GetCategoriesPost(SelectPostLang)
	if err != nil {
		log.Printf("Error in service: %v", err)
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
