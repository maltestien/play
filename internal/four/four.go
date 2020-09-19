package four

import (
	"fmt"
)

type State byte
const (
	Empty State = '-'
	Yellow State = 'Y'
	Red State = 'R'
)

type Board [6][7]State // [rows][columns], row[0] is at the bottom; column[0] is on the left

func NewBoard() *Board {
	return &Board{
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
	}
}

func Players() []byte {
	return []byte{'Y', 'R'}
}

func (b Board) String() string {
	return fmt.Sprintf(
		"| %c | %c | %c | %c | %c | %c | %c |\n" +
		"| %c | %c | %c | %c | %c | %c | %c |\n" +
		"| %c | %c | %c | %c | %c | %c | %c |\n" +
		"| %c | %c | %c | %c | %c | %c | %c |\n" +
		"| %c | %c | %c | %c | %c | %c | %c |\n" +
		"| %c | %c | %c | %c | %c | %c | %c |\n",
		b[5][0], b[5][1], b[5][2], b[5][3], b[5][4], b[5][5], b[5][6],
		b[4][0], b[4][1], b[4][2], b[4][3], b[4][4], b[4][5], b[4][6],
		b[3][0], b[3][1], b[3][2], b[3][3], b[3][4], b[3][5], b[3][6],
		b[2][0], b[2][1], b[2][2], b[2][3], b[2][4], b[2][5], b[2][6],
		b[1][0], b[1][1], b[1][2], b[1][3], b[1][4], b[1][5], b[1][6],
		b[0][0], b[0][1], b[0][2], b[0][3], b[0][4], b[0][5], b[1][6])
}
