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

		if err != nil && err == ticTacToe.ErrPositionNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil && err == ticTacToe.ErrGameAlreadyFinished {
			respond(w, r, http.StatusOK, false)
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
		cells, err := c.ticTacToeService.GetBoard()

		if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		board := ticTacToe.Board{Board: make([]string, 0)}

		// convert rune to string as json has no char/rune like datatype
		for _, cell := range cells {
			board.Board = append(board.Board, string(cell))
		}

		respond(w, r, http.StatusOK, board)
		return
	}
}

func (c *Controller) GetWinner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		winner, err := c.ticTacToeService.GetWinner()

		if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, ticTacToe.Winner{
			Winner: string(winner),
		})
		return
	}
}
