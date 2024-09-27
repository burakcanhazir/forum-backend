package controllers

import (
	"burakforum/models"
	"burakforum/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uuid.New().String()
	err = services.RegisterUser(&user)
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Başarılı yanıt olarak JSON formatında bir mesaj döndür
	response := map[string]string{
		"message": "Register successful",
		"userId":  user.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	fmt.Println("Register Succesful", &user)
}
