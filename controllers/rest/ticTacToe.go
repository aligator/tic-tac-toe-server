package rest

import (
	"github.com/aligator/tic-tac-toe-server/services/ticTacToe"
	"net/http"
)

func (c *Controller) DoMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *Controller) GetFullBoard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		board, err := c.ticTacToeService.GetBoard()

		if err != nil && (err == ticTacToe.ErrPositionNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, board)
		return
	}
}
