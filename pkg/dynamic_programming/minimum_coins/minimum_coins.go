package mc

import (
	"fmt"
	"math"
)

// state transition table
// states[3][9] is the answer, and the minimum coins is 3(it's equal to i of states[i][j])
/*
	0	1	2	3	4	5	6	7	8	9
0	1
1		1		1		1
2			1		1		1		1
3				1		1		1		1
4					1		1		1
5						1		1		1
6							1		1
7								1		1
8									1
9										1
*/
func minimumCoins(coins []int, amount int) int {
	states := make([][]bool, amount+1)
	for i := range states {
		states[i] = make([]bool, amount+1)
	}

	states[0][0] = true

	for i := 1; i <= amount; i++ {
		if states[i-1][amount] {
			break
		}

		for j := 0; j <= amount; j++ {
			for k := range coins {
				if j-coins[k] >= 0 && states[i-1][j-coins[k]] {
					states[i][j] = states[i-1][j-coins[k]]
				}
			}
		}
	}

	printData(states)

	min := -1
	for i := 0; i <= amount; i++ {
		if states[i][amount] {
			min = i
			break
		}
	}

	var path []int
	tmp := amount

	for i := min; i > 0; i-- {
		for k := range coins {
			if tmp-coins[k] >= 0 && states[i-1][tmp-coins[k]] {
				path = append(path, coins[k])
				tmp -= coins[k]
				break
			}
		}
	}
	fmt.Println(path)
	return min
}

// state transition table
func minimumCoins1(coins []int, amount int) int {
	states := make([]bool, amount+1)

	for i := range coins {
		states[coins[i]] = true
	}

	min := -1
	for i := 1; i <= amount; i++ {
		if states[amount] {
			min = i
			break
		}

		for j := amount; j >= 0; j-- {
			if !states[j] {
				continue
			}

			for k := range coins {
				if j+coins[k] <= 9 && !states[j+coins[k]] {
					states[j+coins[k]] = states[j]
				}
			}
		}
	}

	return min
}

// state transition equation
// coins=1+MIN(minimumCoins2(amount-1),minimumCoins2(amount-3),minimumCoins2(amount-5))
func minimumCoins2(coins, mem []int, amount int) int {
	if mem[amount] > 0 {
		return mem[amount]
	}

	for i := range coins {
		if amount == coins[i] {
			return 1
		}
	}

	minCoins := math.MaxInt16
	for i := range coins {
		if amount-coins[i] >= 0 {
			tmp := minimumCoins2(coins, mem, amount-coins[i])
			if tmp < minCoins {
				minCoins = tmp
			}
		}
	}

	mem[amount] = minCoins + 1
	return minCoins + 1
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
