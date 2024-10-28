// Rate Limiting Middleware

package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3) // Saniyede 1 istek, toplamda 3 istek kapasitesi

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "429 - Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
