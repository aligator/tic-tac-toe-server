package ticTacToe

import "github.com/aligator/tic-tac-toe-server/storage"

type Service struct {
	ticTacToe storage.TicTacToe
}

func NewService(t storage.TicTacToe) *Service {
	return &Service{
		t,
	}
}
