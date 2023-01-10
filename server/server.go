package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

func NewServer(router *chi.Mux) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	s.configServer()
	router := s.routes()
	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}
	return nil
}

func (s *Server) configServer() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))
}
