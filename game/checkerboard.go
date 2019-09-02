package checkerboard

import "fmt"

// Size of checkerboard
const (
	RowSize = 6
	ColSize = 7
)

// CheckerBoard .
type CheckerBoard struct {
	board [RowSize][ColSize]int
}

// NewEmptyCheckerBoard creates an empty checkerboard
func NewEmptyCheckerBoard() *CheckerBoard {
	return &CheckerBoard{
		board: [RowSize][ColSize]int{},
	}
}

// Print prints the checkerboard
func (b *CheckerBoard) Print() {
	printPoint := func(point int) {
		if point == 1 {
			fmt.Print("x ")
		} else if point == 2 {
			fmt.Print("o ")
		} else if point == 0 {
			fmt.Print(". ")
		}
	}

	for _, row := range b.board {
		for _, point := range row {
			printPoint(point)
		}
		fmt.Print("\n")
	}
}
