package double11advance

import "fmt"

// items := []int{7, 10, 12, 5, 20, 50, 22, 15, 33}
func double11advance(items []int, w int) {
	states := make([][]bool, len(items))
	capacity := 3 * w

	for i := range items {
		states[i] = make([]bool, capacity+1)
	}

	states[0][0] = true
	if items[0] <= capacity {
		states[0][items[0]] = true
	}

	for i := 1; i < len(items); i++ {
		for j := 0; j <= capacity; j++ {
			if states[i-1][j] {
				// no put the i'th item into capacity
				states[i][j] = states[i-1][j]

				// put the i'th item into capacity, and not exceed capacity
				states[i][j+items[i]] = true
			}
		}
	}

	// j is the first item >=w
	j := w
	// return the maximum in the last layer
	for ; j <= capacity; j++ {
		if states[len(items)-1][j] {
			break
		}
	}

	// not found
	if j == capacity+1 {
		return
	}

	for i := len(items) - 1; i >= 1; i-- {
		if j-items[i] >= 0 && states[i-1][j-items[i]] {
			fmt.Print(items[i], " ")
			j = j - items[i]
		}
	}
	if 0 != j {
		fmt.Print(j)
	}
	fmt.Println()
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
