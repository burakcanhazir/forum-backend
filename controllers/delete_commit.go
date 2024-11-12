package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"burakforum/models"
	"burakforum/services"

	"github.com/gorilla/mux"
)

func DeleteCommit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["postID"]
	commitID := vars["commitID"]
	fmt.Println(commitID)

	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "YETKİLENDİRME HATASI", http.StatusUnauthorized)
		return
	}
	userID := claims.UserID
	err := services.DeleteCommit(userID, postID, commitID)
	if err != nil {
		log.Printf("controllers paketinde hata var: %v", err)
		if err == sql.ErrNoRows {
			http.Error(w, "Kayıt bulunamadı", http.StatusNotFound)
		} else {
			http.Error(w, "İşlem sırasında bir hata oluştu", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent) // Başarı durumunda 204 No Content döndürülür
}
