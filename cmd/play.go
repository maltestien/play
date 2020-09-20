package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"flag"

	"softwaretechnologies/play/internal/four"
)

var reader *bufio.Reader
var totalMoves int
var boardRegistry map[four.Board]four.Player
var interactive *bool

func play(board four.Board, player four.Player, round int) {
	// Determine the available moves
	moves := four.PossibleMoves(board, player)
	if *interactive {
		fmt.Printf("--- R%2d -> Possible Moves: %v; press a key to make the next move...\n", round, moves)
	}

	for _, m := range moves {
		// Make the move
		newBoard, c, r, err := four.MakeMove(board, player, m)
		if err != nil {
			os.Exit(1)
		}
		if *interactive {
			fmt.Printf("    Player %q: Making made: Column %d; Row %d... ", player, c, r)
		}
		totalMoves++

		// fmt.Printf("Resulting board:\n%v\n", newBoard)

		// Check whether we have already come across this Board
		if _, ok := boardRegistry[*newBoard]; ok {
			continue
		} else {
			// If we have never come across this Board, check whether we the current Player has won
			if won, reason := newBoard.PlayerWon(player); won {
				fmt.Printf("Resulting board: %c won (%s).\n%v", player, reason, newBoard)
				boardRegistry[*newBoard] = player
				continue
			} else {
				boardRegistry[*newBoard] = four.Player('-')
			}
		}

		// Ask for the user to hit the RETURN key
		// _, _ = reader.ReadString('\n')

		// If we have not won; play the next round
		play(*newBoard, four.NextPlayer(player), round+1)
	}
}

func main() {
	startTime := time.Now()

	interactive = flag.Bool("interactive", false, "enables interactive mode")
	flag.Parse()

	reader = bufio.NewReader(os.Stdin)
	boardRegistry = make(map[four.Board]four.Player)

	board := four.NewBoard() // Create a new empty board
	player := four.Players()[0] // Determine the first player
	round := 1

	play(board, player, round)

	duration := time.Now().Sub(startTime)

	fmt.Printf("Finished (%v). Total moves: %d; Number of distinct board positions: %d\n", duration, totalMoves, len(boardRegistry))

	var red, yellow int
	for _, v := range boardRegistry {
		if v == 'R' {
			red++
		} else if v == 'Y' {
			yellow++
		}
	}
	fmt.Printf("Games won by Red: %d. Games won by Yellow: %d. %c started\n", red, yellow, player)
}
