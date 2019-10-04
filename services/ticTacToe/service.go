package ticTacToe

import (
	"errors"
	"github.com/aligator/tic-tac-toe-server/storage"
)

var (
	ErrPositionNotFound = errors.New("the requested position could not be found")
)

type Service struct {
	ticTacToe storage.TicTacToe
}

func NewService(t storage.TicTacToe) *Service {
	return &Service{
		t,
	}
}

func (s Service) GetBoard() (Board, error) {
	cells, err := s.ticTacToe.GetBoard()

	if err != nil {
		return Board{}, err
	}

	board := Board{Board: make([]string, 0)}

	// convert rune to string
	for _, cell := range cells {
		board.Board = append(board.Board, string(cell))
	}
	return board, nil
}
