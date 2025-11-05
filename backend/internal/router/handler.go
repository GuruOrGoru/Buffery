package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/guruorgoru/buffery/internal/models"
	"gorm.io/gorm"
)

func handlePing(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			errorMessage := fmt.Sprintf("Error while fetching all users: %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
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
			return
		}
	}
}

func handleCreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestUserBody := new(models.UserRequestBody)
		if err := json.NewDecoder(r.Body).Decode(requestUserBody); err != nil {
			http.Error(w, "Incorrect response pattern", http.StatusBadRequest)
			return
		}
		avatar := ""

		file, handler, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			filepath := fmt.Sprintf("uploads/%s", handler.Filename)
			dst, err := os.Create(filepath)
			if err != nil {
				http.Error(w, "Unable to save avatar", http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, "Error while oploading avatar", http.StatusInternalServerError)
				return
			}
			avatar = filepath
		}

		hashedPassword, err := getHashPassword(requestUserBody.Password)
		if err != nil {
			http.Error(w, "Error while hashing password, again!", http.StatusInternalServerError)
			return
		}

		user := &models.User{
			FullName:     requestUserBody.FullName,
			Email:        requestUserBody.Email,
			PasswordHash: hashedPassword,
			Avatar:       avatar,
		}

		createdUser, err := models.CreateNewUser(db, user)
		if err != nil {
			errorMessage := fmt.Sprintf("Error while inserting new user %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		responseMessage := struct {
			Message models.User `json:"message"`
			Code    string      `json:"code"`
		}{
			*createdUser,
			"200",
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseMessage); err != nil {
			http.Error(w, "Error while encoding output", http.StatusInternalServerError)
			return
		}
	}
}

func handleLoginUser(db *gorm.DB, authToken *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestUserBody := new(models.UserRequestBody)
		if err := json.NewDecoder(r.Body).Decode(requestUserBody); err != nil {
			http.Error(w, "Incorrect response pattern", http.StatusBadRequest)
			return
		}

		userWithThatEmail, err := models.GetByField(db, &models.User{Email: requestUserBody.Email})
		if err != nil {
			http.Error(w, "Error while searching", http.StatusInternalServerError)
			return
		}

		if userWithThatEmail == nil {
			http.Error(w, "User with that email not found", http.StatusNotFound)
			return
		}

		if !checkPassword(userWithThatEmail.PasswordHash, requestUserBody.Password) {
			http.Error(w, "Incorrect password entered", http.StatusBadRequest)
			return
		}

		claims := map[string]any{"id": userWithThatEmail.Id, "email": userWithThatEmail.Email}
		_, tokenStr, err := authToken.Encode(claims)
		if err != nil {
			http.Error(w, "Error while generating token string in login handler", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(tokenStr))
		if err != nil {
			http.Error(w, "Error while writing to the response", http.StatusInternalServerError)
			return
		}
	}
}

func handleGetUserById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userIdInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Please provide correct url params", http.StatusBadRequest)
			return
		}

		_, claims, _ := jwtauth.FromContext(r.Context())
		userIdFromClaims := int(claims["id"].(float64))

		if userIdInt != userIdFromClaims {
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		userWithThatId, err := models.GetByField(db, &models.User{Id: userIdInt})
		if err != nil {
			http.Error(w, "Error while searching", http.StatusInternalServerError)
			return
		}

		if userWithThatId == nil {
			http.Error(w, "User with that id not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(userWithThatId); err != nil {
			http.Error(w, "Error while encoding data", http.StatusInternalServerError)
			return
		}
	}
}
