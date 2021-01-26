package triangle

import (
	"fmt"
	"math"
)

var n = 5

func triangle(data [][]int) int {
	states := make([][]int, n)

	for i := range states {
		states[i] = make([]int, n)
		for j := range states[i] {
			states[i][j] = math.MaxInt16
		}
	}

	states[0][0] = 5

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if -1 == data[i][j] {
				continue
			}

			if j-1 >= 0 {
				states[i][j] = states[i-1][j-1] + data[i][j]
			}

			if -1 != states[i-1][j] && states[i][j] > states[i-1][j]+data[i][j] {
				states[i][j] = states[i-1][j] + data[i][j]
			}

		}
	}
	printData(states)

	min := states[n-1][0]
	for i := 1; i < n; i++ {
		if states[n-1][i] < min {
			min = states[n-1][i]
		}
	}

	return min
}

func triangle1(data [][]int) int {
	states := make([]int, n)

	for i := range states {
		states[i] = math.MaxInt16
	}

	states[0] = 5

	for i := 1; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if -1 == data[i][j] {
				continue
			}
			tmp := math.MaxInt16
			if j-1 >= 0 {
				tmp = states[j-1] + data[i][j]
			}

			if tmp > states[j]+data[i][j] {
				states[j] = states[j] + data[i][j]
			} else {
				states[j] = tmp
			}
		}
	}
	fmt.Println(states)

	min := states[0]
	for i := 1; i < n; i++ {
		if states[i] < min {
			min = states[i]
		}
	}

	return min
}

func printData(data [][]int) {
	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			if math.MaxInt16 != data[row][column] {
				fmt.Print(data[row][column], " ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
