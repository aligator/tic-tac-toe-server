package rest

import (
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (c *Controller) DoMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		position, err := strconv.Atoi(chi.URLParam(r, "position"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		err = c.ticTacToeService.SetPosition(position)

		if err != nil && (err == ticTacToe.ErrPositionNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

func (c *Controller) GetFullBoard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		board, err := c.ticTacToeService.GetBoard()

		if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, board)
		return
	}
}
