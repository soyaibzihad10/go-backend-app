// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/soyaibzihad10/go-backend-app/routes"
)

var logger = logrus.New()

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("❌ Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	logger.Infof("🚀 Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
