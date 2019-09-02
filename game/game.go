package game

import (
	"fmt"
	"math/rand"
)

// GameA defines the 2nd problem
type GameA struct {
	*CheckerBoard
}

// Step places a piece to the checkerboard and checks if it is full
func (a *GameA) Step(piece int) (win bool) {
	col := rand.Intn(ColSize)
	for row := RowSize - 1; row >= 0; row-- {
		if a.board[row][col] == 0 {
			a.board[row][col] = piece
			break
		}
	}
	return a.GameOver()
}

// GameOver when the checkerboard is full
func (a *GameA) GameOver() bool {
	return a.Full()
}

// GameIt simulates one game
func GameIt() {
	game := GameA{
		NewEmptyCheckerBoard(),
	}

	fmt.Println("Before Game:")
	game.Print()

	winner := 0
	for round := 0; winner == 0; round++ {
		player := round%2 + 1
		if game.Step(player) {
			winner = player
		}
	}

	fmt.Println("After Game:")
	game.Print()
}
