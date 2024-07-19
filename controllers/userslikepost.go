package controllers

import (
	"encoding/json"
	"net/http"

	"burakforum/models"
	"burakforum/services"
)

func UsersLikesPost(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID

	post, err := services.UsersLikesPost(userID)
	if err != nil {
		http.Error(w, "POSTLAR GÖRÜNTÜLENİRKEN SORUN OLUŞTU", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
