package controllers

import (
	"encoding/json"
	"net/http"

	"forumbackend/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, "Unable to get users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
