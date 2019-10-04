package rest

import (
	"errors"
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"net/http"
)
import "github.com/go-chi/render"

var (
	ErrInvalidNumberFormat = errors.New("the provided number format is not valid")
	ErrProcessingFailed    = errors.New("failed processing the request")
)

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

func respondError(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	render.JSON(w, r, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}
