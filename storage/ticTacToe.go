package storage

type TicTacToe interface {
	GetCell(position int) (rune, error)
	GetBoard() ([]rune, error)
	SetCell(position int, value rune) error
}
