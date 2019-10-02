package memory

import "errors"

type TicTacToe struct {
	board []rune
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		make([]rune, 9),
	}
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
