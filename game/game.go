package game

import (
	"fmt"
	"math/rand"
)

// Game define a game
type Game interface {
	Step(piece int) bool
	GameOver() bool
	Print()
}

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

// GameB defines the 3nd problem
type GameB struct {
	*CheckerBoard
}

// Step places a piece to the checkerboard and checks if it is full
func (a *GameB) Step(piece int) (win bool) {
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
func (a *GameB) GameOver() bool {
	return a.Full() || a.CheckAll() != 0
}

// GameIt plays the game
func GameIt(game Game) {
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

	if winner != 0 {
		fmt.Println("Winner: Player" + string(winner))
	} else {
		fmt.Println("No Winner")
	}
}
