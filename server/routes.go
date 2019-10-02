package server

import (
	"github.com/go-chi/chi"
)

func (s *Server) mountRoutes() {
	r := chi.NewRouter()

	r.Route("/game", func(r chi.Router) {
		r.Route("/board", func(r chi.Router) {
			r.Get("/", nil)
			r.Route("/{position}", func(r chi.Router) {
				r.Get("/", s.controller.GetBoardPosition())
				r.Put("/", s.controller.DoMove())
			})
		})
	})
}
