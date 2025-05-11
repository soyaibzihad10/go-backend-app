package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Home Route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🏠 Home Page")
}

// About Route
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ℹ️ This is the About Page")
}

// Hello with query param ?name=Soyaib
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "👋 Hello, %s!\n", name)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"message": "This is a JSON response",
	}
	// Encoding the struct to JSON and writing the response
	json.NewEncoder(w).Encode(data)
}

// Path param: /users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)
	userID := urlVars["id"]
	fmt.Println(w, "You requested user ID:\n", userID)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/about", AboutHandler).Methods("GET")
	r.HandleFunc("/hello", HelloHandler).Methods("GET")
	r.HandleFunc("/api/json", JsonHandler).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
