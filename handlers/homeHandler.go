package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Home route accessed")
	fmt.Fprintln(w, "🏠 Welcome to Go Backend")
}
