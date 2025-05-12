// handlers/userHandler.go
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/soyaibzihad10/go-backend-app/db"
	"github.com/soyaibzihad10/go-backend-app/models"
)

var logger = logrus.New()

// Register handler for user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decode JSON request to Go struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error("❌ Error decoding JSON:", err)
		http.Error(w, "❌ Invalid JSON", http.StatusBadRequest)
		return
	}
	// Validate input fields
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" {
		logger.Warn("❌ Name and Email are required")
		http.Error(w, "❌ Name and Email are required", http.StatusBadRequest)
		return
	}

	if !strings.Contains(user.Email, "@") {
		logger.Warn("❌ Invalid email format")
		http.Error(w, "❌ Invalid email format", http.StatusBadRequest)
		return
	}

	// Insert user data into the database
	db.ConnectDatabase()
	conn := db.GetDB()
	fmt.Println(conn)
	_, err = conn.Exec(context.Background(), "INSERT INTO users(name, email) VALUES($1, $2)", user.Name, user.Email)
	if err != nil {
		logger.Error("❌ Error inserting user into database:", err)
		http.Error(w, "❌ Error saving user", http.StatusInternalServerError)
		return
	}

	// Success response
	logger.Info("✅ Registration successful for user:", user.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Registration successful",
		"name":    user.Name,
		"email":   user.Email,
	})
}
