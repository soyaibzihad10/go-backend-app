package middleware

import (
	"net/http"
	"strings"

	"github.com/soyaibzihad10/go-backend-app/auth"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "❌ Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "❌ Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
