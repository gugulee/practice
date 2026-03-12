package lottery

import "fmt"

// decrease decreate []int{3,2,1} to []int{0,0,0} gradually
func decrease(c []int) {
	for i := range c {
		if 0 != c[i] {
			c[i]--
			break
		}
	}
}

// delElements delete target elements from c
func delElements(c []string, target []string) []string {
	newc := c
	for _, t := range target {
		newc = delElement(newc, t)
	}
	return newc
}

// delElement delete target element from c
func delElement(c []string, target string) []string {
	var newc []string
	for i := range c {
		if target == c[i] {
			newc = append(newc, c[0:i]...)
			newc = append(newc, c[i+1:]...)
			return newc
		}
	}
	return nil
}

// in a total of 10 person, the number of third prize is 3,
// the number of second prize is 2,the number of first prize is 1,
// combinations=120*21*5=12600
//
// prize format: int[3,2,1], it means the number of third prize is 3,the number of second prize is 2,the number of first prize is 1
func lottery(origin, rest, result []string, prize []int) {
	// it means that the third prize is over, go to the second prize and so on
	if 0 == prize[0] {
		origin = rest
		prize = prize[1:]
	}

	// it means lottery game is over, print result
	if 0 == len(prize) {
		fmt.Println(result)
		return
	}

	for i := range origin {
		var newResult []string
		newResult = append(newResult, result...) // copy result
		newResult = append(newResult, origin[i])

		var newRest []string
		newRest = delElement(rest, origin[i])

		var newPrize []int
		newPrize = append(newPrize, prize...)
		decrease(newPrize)

		lottery(origin[i+1:], newRest, newResult, newPrize)
	}
}
