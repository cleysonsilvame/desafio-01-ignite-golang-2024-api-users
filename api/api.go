package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHanlder(app application) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handleGetUsers(app))
		r.Post("/users", handleCreateUser(app))
		r.Get("/users/{id}", handleGetUserByID(app))
		r.Patch("/users/{id}", handleUpdateUserByID(app))
		r.Delete("/users/{id}", handleDeleteUserByID(app))
	})

	return r
}
