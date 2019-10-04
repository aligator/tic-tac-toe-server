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

func (s Service) GetPosition(position int) (rune, error) {

	return 'X', nil
}
