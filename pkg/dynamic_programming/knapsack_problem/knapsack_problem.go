package kp

import (
	"fmt"
)

func kp(items []int, w int) (result int) {
	states := make([][]bool, len(items))

	for i := range items {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true
	if items[0] <= w {
		states[0][items[0]] = true
	}

	for i := 1; i < len(items); i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] == true {
				// not put the i'th item into w
				states[i][j] = true

				// put the i'th item into w, and not exceed w
				if j+items[i] <= w {
					states[i][j+items[i]] = true
				}
			}
		}
	}
	printData(states)

	// return the maximum in the last layer
	for i := w; w >= 0; i-- {
		if true == states[len(items)-1][i] {
			return i
		}
	}

	return -1
}

func kp1(items []int, w int) (result int) {
	states := make([]bool, w+1)

	states[0] = true

	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 1; i < len(items); i++ {
		// calculate from end to start, otherwise get the incorrect answer
		for j := w - items[i]; j >= 0; j-- {
			if true == states[j] {
				// put the i'th item into w, and not exceed w
				states[j+items[i]] = true
			}
		}
	}

	for i := w; w >= 0; i-- {
		if true == states[i] {
			return i
		}
	}

	return -1
}

func kp2(items, value []int, w int) (maxW, maxV int) {
	states := make([][]int, len(items))

	for i := range items {
		states[i] = make([]int, w+1)
		for j := range states[i] {
			states[i][j] = -1
		}
	}

	states[0][0] = 0
	if items[0] <= w {
		states[0][items[0]] = value[0]
	}

	for i := 1; i < len(items); i++ {
		for j := 0; j <= w; j++ {
			if -1 != states[i-1][j] {
				// not put the i'th item into w
				if states[i-1][j] > states[i][j] {
					states[i][j] = states[i-1][j]
				}

				// put the i'th item into w, and not exceed w
				if j+items[i] <= w {
					states[i][j+items[i]] = states[i-1][j] + value[i]
				}
			}
		}
	}
	printData1(states)

	maxV = -1
	// return the maximum in the last layer
	for i := 0; i <= w; i++ {
		if states[len(items)-1][i] > maxV {
			maxV = states[len(items)-1][i]
			maxW = i
		}
	}

	return maxW, maxV
}

func kp3(items, value []int, w int) (maxW, maxV int) {
	states := make([]int, w+1)

	for i := range states {
		states[i] = -1
	}

	states[0] = 0
	if items[0] <= w {
		states[items[0]] = value[0]
	}

	for i := 1; i < len(items); i++ {
		// calculate from end to start, otherwise get the incorrect answer
		for j := w - items[i]; j >= 0; j-- {
			if -1 != states[j] && states[j]+value[i] > states[j+items[i]] {
				// put the i'th item into w, and not exceed w
				states[j+items[i]] = states[j] + value[i]
			}
		}
	}

	fmt.Println(states)
	maxV = -1
	// return the maximum in the last layer
	for i := 0; i <= w; i++ {
		if states[i] > maxV {
			maxV = states[i]
			maxW = i
		}
	}

	return maxW, maxV
}

func printData(data [][]bool) {
	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			if true == data[row][column] {
				fmt.Print("1 ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}

func printData1(data [][]int) {
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
