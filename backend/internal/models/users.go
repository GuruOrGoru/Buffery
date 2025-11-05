package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id           int    `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Avatar       string `json:"avatar_url"`
}

type UserRequestBody struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func CreateNewUser(db *gorm.DB, user *User) (*User, error) {
	result := db.Create(user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, errors.New("no rows affected, errors occured")
	}

	return user, nil
}

func GetByField[T any](db *gorm.DB, filter *T) (*T, error) {
	var model T
	result := db.Where(filter).First(&model)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &model, result.Error
}
