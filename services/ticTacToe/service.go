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

func (s Service) GetBoard() ([]rune, error) {
	return s.ticTacToe.GetBoard()
}

func (s Service) SetPosition(position int) error {
	currentPlayer, err := s.ticTacToe.GetCurrentPlayer()

	if err != nil {
		return err
	}

	err = s.ticTacToe.SetCell(position, currentPlayer)

	if err != nil {
		return err
	}

	err = s.ticTacToe.SetCurrentPlayer(getNextPlayer(currentPlayer))

	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetWinner() (rune, error) {
	board, err := s.ticTacToe.GetBoard()

	if err != nil {
		return ' ', err
	}

	checks := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, check := range checks {
		if board[check[0]] == board[check[1]] && board[check[0]] == board[check[2]] {
			return board[check[0]], nil
		}
	}

	return ' ', nil
}

func getNextPlayer(currentPlayer rune) rune {
	if currentPlayer == 'X' {
		return 'O'
	}

	return 'X'
}
