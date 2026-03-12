package get_reward

import "fmt"

var rewards = []int{1, 2, 5, 10}

// GetReward get various combination of reward
func GetReward(total int, result []int) {
	// the condition is satisfied, exit it and print result
	if 0 == total {
		fmt.Println(result)
		return
	}
	// the condition is not satisfied, return
	if total < 0 {
		return
	}

	for _, reward := range rewards {
		var newResult []int
		newResult = append(newResult, result...) // copy result
		newResult = append(newResult, reward)
		GetReward(total-reward, newResult)
	}
}
