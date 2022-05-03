package sudoku

import "fmt"

var found = false

// rol[0][3]=1 represent there is 3 at the row 0
// col[0][4]=1 represent there is 4 at the column 0
// col[0][5]=1 represent there is 4 in the first sudoku(9 in all) from left to right and top to bottom
func sudoku(s, rol, col, grid [][]int, row, column int) {
	if found {
		return
	}

	// found answer
	if row == 9 && column == 0 {
		printData(s)
		found = true
		return
	}

	// there already is a numbner in s[row][column], skip it
	if s[row][column] != 0 {
		if column == 8 {
			sudoku(s, rol, col, grid, row+1, 0)
		} else {
			sudoku(s, rol, col, grid, row, column+1)
		}
		return
	}

	for i := 1; i <= 9; i++ {
		if found {
			return
		}

		// number i does exit at row or column or grid
		if rol[row][i] != 0 || col[column][i] != 0 || grid[column/3+row/3*3][i] != 0 {
			continue
		}

		s[row][column] = i
		rol[row][i] = 1
		col[column][i] = 1
		grid[column/3+row/3*3][i] = 1

		if column == 8 {
			sudoku(s, rol, col, grid, row+1, 0)
		} else {
			sudoku(s, rol, col, grid, row, column+1)
		}

		// rollback
		s[row][column] = 0
		rol[row][i] = 0
		col[column][i] = 0
		grid[column/3+row/3*3][i] = 0
	}
}

func printData(in [][]int) {
	for row := 0; row < len(in); row++ {
		for column := 0; column < len(in[row]); column++ {
			if in[row][column] != 0 {
				fmt.Print(in[row][column], " ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
