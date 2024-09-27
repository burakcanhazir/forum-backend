package controllers

import (
	"burakforum/models"
	"burakforum/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		// Hata mesajını logla
		log.Println("Error fetching user liked posts:", err) // Hata loglama
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
	fmt.Println("SUCCESFUL USERSLİKEPOST")
}
