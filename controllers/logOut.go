package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"burakforum/database"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	// "Bearer " prefixini kaldır
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	Query := "INSERT INTO blacklist (tokenstring) VALUES (?)"
	_, err := database.DB.Exec(Query, tokenString)
	if err != nil {
		log.Printf("error when adding token ")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",                           // Cookie'nin geçerli olduğu path
		Domain:   "http://localhost:8081/login", // Cookie'nin ayarlandığı domain
		MaxAge:   -1,                            // Cookie'yi hemen sil
		Expires:  time.Unix(0, 0),               // Geçmiş bir tarihe ayarla
		HttpOnly: true,                          // Sadece HTTP isteklerinde kullanılabilir
		Secure:   false,                         // Eğer HTTPS kullanmıyorsan, secure bayrağı false olmalı
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged out"))
	fmt.Println("succesfull")
}
