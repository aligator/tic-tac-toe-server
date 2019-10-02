package rest

import (
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"net/http"
)
import "github.com/go-chi/render"

type Controller struct {
	ticTacToeService *ticTacToe.Service
}

func NewController(t *ticTacToe.Service) *Controller {
	controller := Controller{
		t,
	}

	return &controller
}

func respond(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}
