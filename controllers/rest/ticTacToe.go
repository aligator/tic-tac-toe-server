package rest

import (
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (c *Controller) GetBoardPosition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		position, err := strconv.Atoi(chi.URLParam(r, "position"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		value, err := c.ticTacToeService.GetPosition(position)

		if err != nil && (err == ticTacToe.ErrPositionNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, ticTacToe.Position{
			Position: position,
			Value:    string(value),
		})
		return
	}
}

func (c *Controller) DoMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *Controller) GetFullBoard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
