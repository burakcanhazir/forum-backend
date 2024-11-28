package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"forumbackend/models"
	"forumbackend/services"
)

func GetUsersPostsID(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	fmt.Println(userID)

	usersposts, err := services.GetUsersPostsID(userID)
	if err != nil {
		log.Println("CONTROLLERS PAKETİNDE HATA / USERIN POSTLARINI ÇEKERKEN")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersposts)
	fmt.Println("SUCCESFULL GETUSERSPOSTSID")
}
