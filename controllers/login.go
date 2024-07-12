package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"burakforum/middleware"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := middleware.AuthenticateUser(credentials.Name, credentials.Password)
	if err != nil {
		http.Error(w, "Failed to authenticate user: "+err.Error(), http.StatusUnauthorized)
		// http.Redirect(w, r, "/api/v1/register", http.StatusMovedPermanently) eğer kullanıcı bulunamadıysa bunu ekleyecem
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
}
