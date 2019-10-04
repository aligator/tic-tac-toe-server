package storage

type TicTacToe interface {
	GetCurrentPlayer() (rune, error)
	SetCurrentPlayer(nextPlayer rune) error
	GetCell(position int) (rune, error)
	GetBoard() ([]rune, error)
	SetCell(position int, value rune) error
}
