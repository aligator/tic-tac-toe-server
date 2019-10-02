package server

import (
	"github.com/aligator/tic-tac-toe-server/controllers/rest"
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"github.com/aligator/tic-tac-toe-server/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Server struct {
	router *chi.Mux
	http   *http.Server

	ticTacToe storage.TicTacToe

	ticTacToeService *ticTacToe.Service

	controller *rest.Controller
}

func New() (*Server, error) {
	s := Server{}
	var err error

	s.router = newRouter()
	s.http = newHTTPServer(s.router, ":5000")

	s.controller = rest.NewController(s.ticTacToeService)

	s.mountRoutes()

	return &s, err
}

func newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)
	return router
}

func newHTTPServer(handler http.Handler, addr string) *http.Server {
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &server
}
