package config

import (
	"errors"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type App struct {
	DB        *gorm.DB
	AuthToken *jwtauth.JWTAuth
}

func InitApp(gormDB *gorm.DB, jwtSecret string) *App {
	app := App{
		DB:        gormDB,
		AuthToken: GenerateAuthToken(jwtSecret),
	}

	return &app
}

func GetPort() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	portstr := os.Getenv("PORT")
	if portstr == "" {
		return "", errors.New("ENV Port not set in .env")
	}
	return portstr, nil
}

func GetHost() (string, error) {
	host := os.Getenv("HOST")
	if host == "" {
		return "", errors.New("ENV Host not set in .env")
	}
	return host, nil
}

func GetDbURL() (string, error) {
	url := os.Getenv("DB_URL")
	if url == "" {
		return "", errors.New("ENV DB_URL not set in .env")
	}
	return url, nil
}

func GetJwtKey() (string, error) {
	url := os.Getenv("JWT_SECRET")
	if url == "" {
		return "", errors.New("ENV JWT_SECRET not set in .env")
	}
	return url, nil
}

func GenerateAuthToken(jwtSecret string) *jwtauth.JWTAuth {
	tokenOfAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)
	return tokenOfAuth
}
