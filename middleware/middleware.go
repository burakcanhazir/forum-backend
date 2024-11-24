package middleware

import (
	"context"
	"log"
	"net/http"
	"os"

	"burakforum/database"

	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

// AuthMiddleware yetkilendirme kontrolü yapan middleware fonksiyonu
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cookie'den token'ı al
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Token not found in cookies", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		// Token'ı blacklist'e göre kontrol et
		rows, err := database.DB.Query("SELECT tokenstring FROM blacklist")
		if err != nil {
			log.Fatalf("Failed to execute query: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var isblocked string
			if err := rows.Scan(&isblocked); err != nil {
				log.Fatalf("Failed to scan row: %v", err)
			}
			if isblocked == tokenString {
				http.Error(w, "Session expired", http.StatusBadRequest)
				return
			}
		}

		// Token'ı doğrula
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Claims'leri context'e ekle
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
