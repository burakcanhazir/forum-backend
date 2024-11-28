package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forumbackend/middleware"
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
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteLaxMode, // Lax modunu kullan
		Secure:   false,                // HTTP üzerinde true yapma
		HttpOnly: true,                 // Güvenlik için doğru
	})
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Bunu ekle

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: token})
	fmt.Println("login succesful", token)
}
