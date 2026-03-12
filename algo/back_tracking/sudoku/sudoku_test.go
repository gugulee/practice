package sudoku

import (
	"fmt"
	"testing"
)

func generateTestData() [][]int {
	s := make([][]int, 9)
	for i := 0; i < 9; i++ {
		s[i] = make([]int, 9)
	}

	// the first line
	s[0][2] = 5
	s[0][3] = 3

	// the second line
	s[1][0] = 8
	s[1][7] = 2

	// the third line
	s[2][1] = 7
	s[2][4] = 1
	s[2][6] = 5

	// the fourth line
	s[3][0] = 4
	s[3][5] = 5
	s[3][6] = 3

	// the fifth line
	s[4][1] = 1
	s[4][4] = 7
	s[4][8] = 6

	// the sixth line
	s[5][2] = 3
	s[5][3] = 2
	s[5][7] = 8

	// the seventh line
	s[6][1] = 6
	s[6][3] = 5
	s[6][8] = 9

	// the eighth line
	s[7][2] = 4
	s[7][7] = 3

	// the ninth line
	s[8][5] = 9
	s[8][6] = 7

	return s
}

func Test_sudoku(t *testing.T) {
	s := generateTestData()
	rol := make([][]int, 9)
	col := make([][]int, 9)
	grid := make([][]int, 9)

	for i := 0; i < 9; i++ {
		rol[i] = make([]int, 10)
		col[i] = make([]int, 10)
		grid[i] = make([]int, 10)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if 0 == s[i][j] {
				continue
			}
			rol[i][s[i][j]] = 1
			col[j][s[i][j]] = 1
			grid[j/3+i/3*3][s[i][j]] = 1
		}
	}
	printData(s)
	fmt.Println("----------------------------")

	// start from s[0][0]
	sudoku(s, rol, col, grid, 0, 0)
}
