package four

import (
	"fmt"
	"os"
	// "errors"
)

type Player byte
type State byte
const (
	Empty State = '-'
	Yellow State = 'Y'
	Red State = 'R'
)
type Move int // the column index of the column to toss into

const rows = 4
const columns = 7

type Board [rows][columns]State // [rows][columns], row[0] is at the bottom; column[0] is on the left

func NewBoard() Board {
	var b Board

	for r := rows-1; r>=0; r-- {
		for c := 0; c<columns; c++ {
			b[r][c] = Empty
		}
	}

	return b
}

func Players() []Player {
	return []Player{'Y', 'R'}
}

func NextPlayer(p Player) Player {
	if p == 'Y' {
		return 'R'
	} else if p == 'R' {
		return 'Y'
	} else {
		fmt.Fprintf(os.Stderr, "Error: Unable to determined next player for player %q\n", p)
		os.Exit(1)
		return '-'
	}
}

func PossibleMoves(board Board, player Player) []Move {
	var result []Move
	for i := 0; i<columns; i++ {
		if board[rows-1][i] == '-' {
			result = append(result, Move(i))
		}
	}
	return result
}

func MakeMove(b Board, p Player, m Move) (board *Board, column int, row int, Error error) {
	newBoard := b.Copy()
	column = int(m)

	if b[rows-1][m] != Empty {
		return nil, 0, 0, fmt.Errorf("MakeMove: impossible to make move %d; column full", int(m))
	}

	for r := rows-1; r>=0; r-- {
		if b[r][m] != Empty {
			(*newBoard)[r+1][m] = State(p)
			row = r+1
			break
		} else if r == 0 {
			(*newBoard)[r][m] = State(p)
			row = r
			break
		}
	}

	return newBoard, column, row, nil
}

func (b Board) String() string {
	var result string

	for r := rows-1; r>=0; r-- {
		result += "    |"
		for c := 0; c<columns; c++ {
			result += " " + string(b[r][c]) + " |"
		}

		result += "\n"
	}

	return result
}

func (b Board) PlayerWon(p Player) (bool, string) {
	// Check for horizontal victory
	if columns>=4 {
		for r := 0; r<rows; r++ {
			for c := 0; c<=columns-4; c++ {
				if b[r][c] == State(p) && b[r][c+1] == State(p) && b[r][c+2] == State(p) && b[r][c+3] == State(p) {
					return true, "horizontal"
				}
			}
		}
	}

	// Check for vertical victory
	if rows>=4 {
		for r := 0; r<=rows-4; r++ {
			for c := columns-1; c>=0; c-- {
				if b[r][c] == State(p) && b[r+1][c] == State(p) && b[r+2][c] == State(p) && b[r+3][c] == State(p) {
					return true, "vertical"
				}
			}
		}
	}

	// Check for diagonal victory
	if columns >= 4 && rows>=4{
		for r := 0; r<=rows-4; r++ {
			for c := 0; c<=columns-4; c++ {
				if b[r][c] == State(p) && b[r+1][c+1] == State(p) && b[r+2][c+2] == State(p) && b[r+3][c+3] == State(p) {
					return true, "diagonally, bottom-left to top-right" // bottom-left to top-right
				}
				if b[r+3][c] == State(p) && b[r+2][c+1] == State(p) && b[r+1][c+2] == State(p) && b[r][c+3] == State(p) {
					return true, "diagonally, top-left to bottom-right" // top-left to bottom-right
				}
			}
		}
	}

	return false, ""
}

func (b Board) Copy() *Board {
	var newBoard Board

	for row := range b {
		newBoard[row] = b[row]
	}

	return &newBoard
}
