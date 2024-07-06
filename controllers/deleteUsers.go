package controllers

import (
	"encoding/json"
	"net/http"

	"burakforum/models"
	"burakforum/services"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "AYIKLAMA YAPILIRKEN HATA OLUŞTU", http.StatusBadRequest)
		return
	}
	err = services.DeleteUsers(&user)
	if err != nil {
		http.Error(w, "silinemedi-controllers paketine bak", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
