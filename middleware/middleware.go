package middleware

import (
	"context"
	"net/http"
	"strings"
)

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

		// Token'ı doğrula
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Token geçerliyse isteği bir sonraki handler'a geçir
		// İsteğin context'ine claims ekleyebiliriz (opsiyonel)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userClaims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
