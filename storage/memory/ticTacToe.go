package memory

import "errors"

type TicTacToe struct {
	currentPlayer rune
	board         []rune
}

func NewTicTacToe() *TicTacToe {
	ticTacToe := TicTacToe{
		currentPlayer: 'X',
		board:         make([]rune, 9),
	}

	for i, _ := range ticTacToe.board {
		ticTacToe.board[i] = ' '
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
		return ' ', errors.New("requested position does not exist")
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
