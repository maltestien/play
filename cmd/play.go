package main

import (
	"fmt"
	"os"
	"bufio"

	"softwaretechnologies/play/internal/four"
)

var reader *bufio.Reader

func play(board four.Board, player four.Player, round int) {
	// Determine the available moves
	moves := four.PossibleMoves(board, player)
	fmt.Printf("--- R%2d -> Possible Moves: %v; press a key to make the next move...\n", round, moves)

	for _, m := range moves {
		// Make the move
		newBoard, c, r, err := four.MakeMove(board, player, m)
		fmt.Printf("    Player %q: Move: Column %d; Row %d\n", player, c, r)
		if err != nil {
			os.Exit(1)
		}

		// Check whether we have won
		// ...

		// If we have not won; play the next round
		fmt.Printf("%v\n", newBoard)
		_, _ = reader.ReadString('\n')
		play(*newBoard, four.NextPlayer(player), round+1)
	}

	// fmt.Println("Returning...")
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	board := four.NewBoard() // Create a new empty board
	player := four.Players()[0] // Determine the first player
	round := 1

	play(board, player, round)
}
