package md

import (
	"fmt"
	"math"

	"github.com/practice/pkg/tools"
)

func minimumDistance(data [][]int, endRow, endColumn int) int {
	minDist := make([][]int, len(data))
	for i := range minDist {
		minDist[i] = make([]int, len(minDist))
	}

	// initialize the first row
	sum := 0
	for j := range minDist[0] {
		sum += data[0][j]
		minDist[0][j] = sum
	}

	// initialize the first column
	sum = 0
	for i := 0; i < len(data); i++ {
		sum += data[i][0]
		minDist[i][0] = sum
	}

	for i := 1; i < len(data); i++ {
		for j := 1; j < len(data); j++ {
			minDist[i][j] = data[i][j] + tools.Min(minDist[i][j-1], minDist[i-1][j])
		}
	}

	printData(minDist)
	return minDist[endRow][endColumn]
}

func minimumDistance1(data, mem [][]int, i, j int) int {
	if 0 == i && 0 == j {
		return data[i][j]
	}

	if mem[i][j] > 0 {
		return mem[i][j]
	}

	minLeft := math.MaxInt16
	if j-1 >= 0 {
		minLeft = minimumDistance1(data, mem, i, j-1)
	}

	minUp := math.MaxInt16
	if i-1 >= 0 {
		minUp = minimumDistance1(data, mem, i-1, j)
	}

	currentMinDist := data[i][j] + tools.Min(minLeft, minUp)
	mem[i][j] = currentMinDist
	return currentMinDist
}

func printData(data [][]int) {
	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			if -1 != data[row][column] {
				fmt.Print(data[row][column], " ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
