package controllers

import (
	"encoding/json"
	"net/http"

	"forumbackend/services" // services paketini import et
)

// getPosts HTTP işleyicisi
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// services paketindeki GetPosts fonksiyonunu çağır
	posts, err := services.GetPosts()
	if err != nil {
		http.Error(w, "Unable to get posts", http.StatusInternalServerError)
		return
	}

	// JSON formatında postları yanıtla
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
