package v2

import "fmt"

var count = 0

func eightQueens(in []int, row int) {
	if row == 8 {
		count++
		printQueues(in)
		fmt.Println("------------------------------------")
		return
	}
	
	for column := 0; column < 8; column++ {
		if isOK(in, row, column) {
			in[row] = column
			eightQueens(in, row+1)
		}
	}
}

func isOK(in []int, row, column int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		// check if there is a queue in the column line
		if in[i] == column {
			return false
		}

		// check if there is a queue in the upper left diagonal line
		if leftup >= 0 && in[i] == leftup {
			return false
		}

		// check if there is a queue in the upper right diagonal line
		if rightup < 8 && in[i] == rightup {
			return false
		}

		leftup, rightup = leftup-1, rightup+1
	}

	return true
}

func printQueues(in []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if in[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
