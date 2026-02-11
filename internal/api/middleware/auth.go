package middleware

import (
	"net/http"
	"strings"

	"reservation-system/internal/infrastructure/jwt"
	"reservation-system/pkg/response"
)

type AuthClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Unauthorized(w, "Authorization header required")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Unauthorized(w, "Bearer token required")
			return
		}

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(w, "Invalid token")
			return
		}

		r.Header.Set("X-User-ID", string(rune(claims.UserID)))
		r.Header.Set("X-User-Email", claims.Email)

		next.ServeHTTP(w, r)
	}
}
