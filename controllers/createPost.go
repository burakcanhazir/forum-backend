package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"burakforum/models"
	"burakforum/services"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// JSON verilerini çözümle
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = services.CreatePost(&post)
	if err != nil {
		http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
