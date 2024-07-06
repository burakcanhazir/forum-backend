package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"burakforum/models"
	"burakforum/services"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// İstek gövdesinden JSON verisini oku
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Yeni kullanıcıyı veritabanına ekle
	err = services.CreateUser(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User with the same name already exists", http.StatusConflict)
		} else {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
		}
		return
	}

	// Başarılı yanıt döndür
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
