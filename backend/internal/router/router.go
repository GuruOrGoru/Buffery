package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func GetServer(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", handlePing(db))
	return r
}
