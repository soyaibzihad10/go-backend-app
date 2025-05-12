package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/soyaibzihad10/go-backend-app/auth"
	"github.com/soyaibzihad10/go-backend-app/db"
)

type LoginRequest struct {
	Email string `json:"email"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if !strings.Contains(req.Email, "@") {
		http.Error(w, "❌ Invalid email", http.StatusBadRequest)
		return
	}

	conn := db.GetDB()
	var email string
	err := conn.QueryRow(r.Context(), "SELECT email FROM users WHERE email=$1", req.Email).Scan(&email)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(email)
	if err != nil {
		http.Error(w, "❌ Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
