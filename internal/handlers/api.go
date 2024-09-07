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

		// Auth Middleware
		router.Use(middleware.Authorization)

		// Applicant routes
		router.Get("/applicants", GetApplicants)
		router.Post("/applicants", PostApplicants)
	})

}
