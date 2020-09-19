package main

import (
	"fmt"

	"softwaretechnologies/play/internal/four"
)

func main() {
	// Create a new empty board
	board := four.NewBoard()

	// Get the available players
	players := four.Players()
	round := 1

	for {
		player := players[(round%len(players))]
		fmt.Printf("Round %d; Player %q\n", round, player)

		// Get the available moves
		// ...

		

		// Check whether we have won
		// ..

		// Next round
		round++
	}

	fmt.Printf("%v\n", (*board))
}
