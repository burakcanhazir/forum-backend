package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"burakforum/models"
	"burakforum/services"

	"github.com/gorilla/mux"
)

func CreateCommit(w http.ResponseWriter, r *http.Request) {
	// postID alındı
	vars := mux.Vars(r)
	postID := vars["id"]

	// token sahibinin kullanıcı bilgileri alındı
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}
	var Commit models.Commit

	err := json.NewDecoder(r.Body).Decode(&Commit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("json ayıklamada hata")
		return
	}
	Commit.UserID = claims.UserID
	Commit.PostID = postID

	err = services.CreateCommit(&Commit)
	if err != nil {
		http.Error(w, "commit gönderildi control edilemedi: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Commit)
}
