package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

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
