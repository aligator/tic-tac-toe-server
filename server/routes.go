package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (s *Server) mountRoutes() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Route("/game", func(r chi.Router) {
		r.Get("/winner", s.controller.GetWinner())
		r.Route("/board", func(r chi.Router) {
			r.Get("/", s.controller.GetFullBoard())
			r.Route("/{position}", func(r chi.Router) {
				r.Put("/", s.controller.DoMove())
			})
		})
	})

	s.router.Mount("/", r)
}
