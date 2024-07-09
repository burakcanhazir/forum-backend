package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"burakforum/services"

	"github.com/gorilla/mux"
)

func GetUsersPostsID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	usersposts, err := services.GetUsersPostsID(userID)
	if err != nil {
		log.Println("CONTROLLERS PAKETİNDE HATA / USERIN POSTLARINI ÇEKERKEN")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersposts)
}
