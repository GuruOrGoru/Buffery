package router

import (
	"encoding/json"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	responseMessage := struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}{
		"Pong",
		"200",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseMessage); err != nil {
		http.Error(w, "Error while encoding output", http.StatusInternalServerError)
	}
}
