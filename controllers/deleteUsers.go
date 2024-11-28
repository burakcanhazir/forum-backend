package controllers

import (
	"fmt"
	"net/http"

	"forumbackend/models"
	"forumbackend/services"
)

const UserClaimsKey = "userClaims"

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	// Context'ten kullanıcı kimliğini al
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}
	fmt.Println(claims.UserID)

	// Kullanıcı kimliği eşleşiyor mu kontrol et
	err := services.DeleteUsers(claims.UserID)
	if err != nil {
		http.Error(w, "Silinemedi - controllers paketine bak", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}
