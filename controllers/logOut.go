package controllers

import (
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
		Name:    "token",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged out"))
}
