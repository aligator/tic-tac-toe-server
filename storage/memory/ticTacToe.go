package memory

import (
	"errors"
	"github.com/aligator/tic-tac-toe-server/constants"
)

type TicTacToe struct {
	currentPlayer rune
	board         []rune
}

func NewTicTacToe() *TicTacToe {
	ticTacToe := TicTacToe{
		currentPlayer: constants.PLAYER1_CELL,
		board:         make([]rune, 9),
	}

	for i, _ := range ticTacToe.board {
		ticTacToe.board[i] = constants.EMPTY_CELL
	}

	return &ticTacToe
}

func (t TicTacToe) GetCurrentPlayer() (rune, error) {
	return t.currentPlayer, nil
}

func (t *TicTacToe) SetCurrentPlayer(nextPlayer rune) error {
	t.currentPlayer = nextPlayer

	return nil
}

func (t TicTacToe) GetCell(position int) (rune, error) {
	if len(t.board) <= position {
		return constants.EMPTY_CELL, errors.New("requested position does not exist")
	}

	cell := t.board[position]
	return cell, nil
}

func (t *TicTacToe) SetCell(position int, value rune) error {
	if len(t.board) <= position {
		return errors.New("requested position does not exist")
	}

	t.board[position] = value
	return nil
}

func (t *TicTacToe) GetBoard() ([]rune, error) {
	return t.board, nil
}
