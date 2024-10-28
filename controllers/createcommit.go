package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"burakforum/models"
	"burakforum/services"

	"github.com/gorilla/mux"
)

func CreateCommit(w http.ResponseWriter, r *http.Request) {
	// postID alındı
	fmt.Println("geldi yorum isteği")
	vars := mux.Vars(r)
	postID := vars["id"]

	err := services.TruePostID(postID)
	if err != nil {
		http.Error(w, "POST ID YALNIŞ", http.StatusBadRequest)
		return
	}

	// token sahibinin kullanıcı bilgileri alındı
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	fmt.Println(claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}
	var Commit models.Commit

	err = json.NewDecoder(r.Body).Decode(&Commit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("json ayıklamada hata")
		return
	}

	Commit.UserID = claims.UserID
	Commit.PostID = postID
	fmt.Printf("mesaj geldi", Commit.Content)

	err = services.CreateCommit(&Commit)
	if err != nil {
		http.Error(w, "commit gönderildi control edilemedi: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("evvet oldu")
	json.NewEncoder(w).Encode(Commit)
}
