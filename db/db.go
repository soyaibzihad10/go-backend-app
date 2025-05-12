// db/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

// ConnectDatabase sets up the PostgreSQL database connection
func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Get DB credentials from environment variables
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		"5432", // default PostgreSQL port
		os.Getenv("DB_NAME"),
	)

	// Connect to the database
	Conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("❌ Unable to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
}

// GetDB returns the database connection
func GetDB() *pgx.Conn {
	return Conn
}
