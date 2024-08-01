package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"burakforum/database"

	"github.com/dgrijalva/jwt-go"
)

// Your secret key
var jwtKey = []byte("my_secret_key")

type UserClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// AuthMiddleware yetkilendirme kontrolü yapan middleware fonksiyonu
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization başlığını al
		tokenString := r.Header.Get("Authorization")

		// Başlık yoksa yetkilendirme hatası döndür
		if tokenString == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// "Bearer " prefix'ini kaldır
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// OTURUM SÜRESİ BİTMİŞSE NEXT HANDLER OLMAYACAK
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
				http.Error(w, "oturum süresi bitmiş", http.StatusBadRequest)
				return
			}

		}

		// Token'ı doğrula
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Token geçerliyse isteği bir sonraki handler'a geçir
		// İsteğin context'ine claims ekleyebiliriz
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userClaims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
