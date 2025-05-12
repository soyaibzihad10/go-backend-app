package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soyaibzihad10/go-backend-app/handlers"
	"github.com/soyaibzihad10/go-backend-app/middleware"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ Welcome to protected route"))
	})).Methods("GET")
}
