package controllers

import (
	"burakforum/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := services.GetUserID(userID)
	if err != nil {
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// JSON formatında yanıtı döndür
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
