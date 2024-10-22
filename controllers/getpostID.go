package controllers

import (
	"burakforum/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPostID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("burası controllers back")
	vars := mux.Vars(r)
	postID := vars["id"]
	fmt.Println(postID)

	check, err := services.GetpostID(postID)
	if err != nil {
		fmt.Println("CONTROLLERS PAKETİNDE HATA VAR. GETPOSTID")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if check == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(check)
}
