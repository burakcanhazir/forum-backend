package controllers

import (
	"encoding/json"
	"net/http"

	"burakforum/services"

	"github.com/gorilla/mux"
)

func GetPostCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CatID := vars["id"]

	postCategories, err := services.GetPostCategories(CatID)
	if err != nil {
		http.Error(w, "geçersiz kategori: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if postCategories == nil {
		http.Error(w, "kategori bulunamadı", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postCategories)
}
