package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", handlePing)
	return r
}
