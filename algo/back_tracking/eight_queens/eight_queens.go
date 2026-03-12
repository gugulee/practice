package eightqueens

import "fmt"

var count = 0

func eightQueens(in [][]int, row int) {
	if row == 8 {
		count++
		printQueues(in)
		fmt.Println("------------------------------------")
		return
	}

	for column := 0; column < 8; column++ {
		newin := make([][]int, len(in))
		for i := range in {
			newin[i] = make([]int, len(in[i]))
			copy(newin[i], in[i])
		}

		if isOK(newin, row, column) {
			newin[row][column] = 1
			eightQueens(newin, row+1)
		}
	}
}

// isOK return true if find the empty position
func isOK(in [][]int, row, column int) bool {
	// check if there is a queue in the column line
	for i := row - 1; i >= 0; i-- {
		if 1 == in[i][column] {
			return false
		}
	}

	// check if there is a queue in the upper left diagonal line
	for x, y := row-1, column-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if 1 == in[x][y] {
			return false
		}
	}

	// check if there is a queue in the upper right diagonal line
	for x, y := row-1, column+1; x >= 0 && y < 8; x, y = x-1, y+1 {
		if 1 == in[x][y] {
			return false
		}
	}

	return true
}

// // isOK return true if find the empty position
// func isOK(in [][]int, row, column int) bool {
// 	// check if there is a queue in the column line
// 	for i := row - 1; i >= 0; i-- {
// 		if 1 == in[i][column] {
// 			return false
// 		}
// 	}

// 	for i := row + 1; i < 8; i++ {
// 		if 1 == in[i][column] {
// 			return false
// 		}
// 	}

// 	// check if there is a queue in the row line
// 	for i := column - 1; i >= 0; i-- {
// 		if 1 == in[row][i] {
// 			return false
// 		}
// 	}

// 	for i := column + 1; i < 8; i++ {
// 		if 1 == in[row][i] {
// 			return false
// 		}
// 	}

// 	// check if there is a queue in the left diagonal line
// 	for x, y := row-1, column-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
// 		if 1 == in[x][y] {
// 			return false
// 		}
// 	}

// 	for x, y := row+1, column+1; x < 8 && y < 8; x, y = x+1, y+1 {
// 		if 1 == in[x][y] {
// 			return false
// 		}
// 	}

// 	// check if there is a queue in the right diagonal line
// 	for x, y := row-1, column+1; x >= 0 && y < 8; x, y = x-1, y+1 {
// 		if 1 == in[x][y] {
// 			return false
// 		}
// 	}

// 	for x, y := row+1, column-1; x < 8 && y >= 0; x, y = x+1, y-1 {
// 		if 1 == in[x][y] {
// 			return false
// 		}
// 	}

// 	return true
// }

func printQueues(in [][]int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if in[row][column] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
