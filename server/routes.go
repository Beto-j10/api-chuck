package server

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() *chi.Mux {
	r := s.router
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/chuck", s.getHandler())

	})
	return r
}
