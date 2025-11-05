package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/guruorgoru/buffery/internal/config"
)

func GetServer(app *config.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", handlePing(app.DB))
	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", handleCreateUser(app.DB))
		r.Post("/login", handleLoginUser(app.DB, app.AuthToken))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(app.AuthToken))
			r.Use(jwtauth.Authenticator(app.AuthToken))
			
			r.Get("/{id}", handleGetUserById(app.DB))
		})
	})
	return r
}
