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

const rows = 6
const columns = 7

type Board [rows][columns]State // [rows][columns], row[0] is at the bottom; column[0] is on the left

func NewBoard() Board {
	return Board{
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
	}
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

func (b Board) Copy() *Board {
	var newBoard Board

	for row := range b {
		newBoard[row] = b[row]
	}

	return &newBoard
}
