package controllers

import (
	"forumbackend/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCommits(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PostID := vars["id"]

	// Sadece tek bir dönen değeri alıyoruz.
	Commits := services.GetCommits(PostID)

	// JSON formatında postları yanıtla
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Commits))
}
