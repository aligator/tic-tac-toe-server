package server

import (
	"context"
	"github.com/aligator/tic-tac-toe-server/controllers/rest"
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"github.com/aligator/tic-tac-toe-server/storage"
	"github.com/aligator/tic-tac-toe-server/storage/memory"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	s.ticTacToe = memory.NewTicTacToe()

	s.ticTacToeService = ticTacToe.NewService(s.ticTacToe)

	s.controller = rest.NewController(s.ticTacToeService)

	s.mountRoutes()

	return &s, err
}

func (s *Server) Run() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		if err := s.http.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		log.Println(err)
	}
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
