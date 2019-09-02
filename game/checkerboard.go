package game

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

// Full checks if the checkerboard is full
func (b *CheckerBoard) Full() bool {
	for _, row := range b.board {
		for _, point := range row {
			if point == 0 {
				return false
			}
		}
	}
	return true
}

// CheckAll checks if the checkerboard is full
func (b *CheckerBoard) CheckAll() int {
	// check returns if point is in a winning pattern
	check := func(row, col, point int) int {

		// check row
		count := 1
		for j := col + 1; j <= col+3 && j <= ColSize-1; j++ {
			if point == b.board[row][j] {
				count++
			} else {
				break
			}
		}

		if count == 4 {
			return point
		}

		// check col
		count = 1
		for i := row + 1; i <= row+3 && i <= RowSize-1; i++ {
			if point == b.board[i][col] {
				count++
			} else {
				break
			}
		}

		if count == 4 {
			return point
		}

		// check down right
		count = 1
		for i, j := row+1, col+1; i <= row+3 && i <= RowSize-1 && j <= col+3 && j <= ColSize-1; i, j = i+1, j+1 {
			if point == b.board[i][j] {
				count++
			} else {
				break
			}
		}

		if count == 4 {
			return point
		}

		// check down left
		count = 1
		for i, j := row+1, col-1; i <= row+3 && i <= RowSize-1 && j >= col-3 && j >= 0; i, j = i+1, j-1 {
			if point == b.board[i][j] {
				count++
			} else {
				break
			}
		}

		if count == 4 {
			return point
		}

		return 0
	}

	for r, row := range b.board {
		for c, point := range row {
			win := check(r, c, point)
			if win != 0 {
				return win
			}
		}
	}
	return 0
}
