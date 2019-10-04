package server

import (
	"github.com/go-chi/chi"
)

func (s *Server) mountRoutes() {
	r := chi.NewRouter()

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
