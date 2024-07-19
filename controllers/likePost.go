package controllers

import (
	"log"
	"net/http"

	"burakforum/models"
	"burakforum/services"

	"github.com/gorilla/mux"
)

func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["post_id"]

	// Kullanıcının kimlik doğrulamasını kontrol et
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}
	userID := claims.UserID

	// Beğeni ekle veya çıkar
	err := services.LikePost(postID, userID)
	if err != nil {
		log.Printf("Error in LikePost: %v", err)
		http.Error(w, "Beğeni güncellemesinde hata", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Başarıyla tamamlandığında HTTP 200 döner
}
