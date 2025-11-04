package models

import "gorm.io/gorm"

type User struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Avatar       string `json:"avatar_url"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}
