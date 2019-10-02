package storage

type TicTacToe interface {
	GetCell(position uint) (rune, error)
	SetCell(position uint, value rune) error
}
