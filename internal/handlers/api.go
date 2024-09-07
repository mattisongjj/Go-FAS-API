package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/mattisongjj/Go-FAS-API/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/api", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/applicants", GetApplicants)
	})

}
