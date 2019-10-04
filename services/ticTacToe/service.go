package ticTacToe

import (
	"errors"
	"github.com/aligator/tic-tac-toe-server/constants"
	"github.com/aligator/tic-tac-toe-server/storage"
)

var (
	ErrPositionNotFound    = errors.New("the requested position could not be found")
	ErrGameAlreadyFinished = errors.New("the game is already finished")
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
	// first check if game is not already finished
	winner, err := s.GetWinner()

	if err != nil {
		return err
	}

	if winner != constants.EMPTY_CELL {
		return ErrGameAlreadyFinished
	}

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
		return constants.EMPTY_CELL, err
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

	return constants.EMPTY_CELL, nil
}

func getNextPlayer(currentPlayer rune) rune {
	if currentPlayer == constants.PLAYER1_CELL {
		return constants.PLAYER2_Cell
	}

	return constants.PLAYER1_CELL
}
