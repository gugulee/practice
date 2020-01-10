package get_min_money

import "math"

var moneyType = []int{2, 3, 7}

const MaxCount = math.MaxUint32

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinMoney(total int) int {
	if 0 == total {
		return 0
	}
	minCount := MaxCount
	for _, money := range moneyType {
		if total-money < 0 {
			minCount = min(MaxCount, minCount)
		} else {
			minCount = min(getMinMoney(total-money)+1, minCount)
		}
	}
	return minCount
}

func getMinMoney1(total int) int {
	if 0 == total {
		return 0
	}

	ch := make(chan int, len(moneyType))
	for _, money := range moneyType {
		go func(m int) {
			if total-m < 0 {
				ch <- MaxCount
			} else {
				ch <- getMinMoney(total-m) + 1
			}

		}(money)
	}

	minCount := MaxCount
	for range moneyType {
		minCount = min(<-ch, minCount)
	}

	return minCount
}

func getMinMoney2(total int) int {
	minCount := make(map[int]int)
	for i := 1; i <= total; i++ {
		minCount[i] = MaxCount
	}

	for i := 1; i <= total; i++ {
		for _, money := range moneyType {
			if i-money < 0 {
				minCount[i] = min(MaxCount, minCount[i])
			} else {
				minCount[i] = min(minCount[i-money]+1, minCount[i])
			}
		}
	}

	return minCount[total]
}
