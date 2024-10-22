package controllers

import (
	"burakforum/models"
	"burakforum/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Context'ten kullanıcı kimliğini al
	claims, ok := r.Context().Value(UserClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Yetkilendirme hatası", http.StatusUnauthorized)
		return
	}

	// JSON verilerini çözümle
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.UserID = claims.UserID
	post.ID = uuid.New().String()

	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = services.CreatePost(&post)
	if err != nil {
		http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(post)
	log.Printf("Post Title: %s, Content: %s, Categories: %v", post.Title, post.Content, post.Category)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
