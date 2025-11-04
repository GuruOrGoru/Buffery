package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guruorgoru/buffery/internal/models"
	"gorm.io/gorm"
)

func handlePing(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			errorMessage := fmt.Sprintf("Error while fetching all users: %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
		}
		responseMessage := struct {
			Message []models.User `json:"message"`
			Code    string        `json:"code"`
		}{
			users,
			"200",
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseMessage); err != nil {
			http.Error(w, "Error while encoding output", http.StatusInternalServerError)
		}
	}
}
