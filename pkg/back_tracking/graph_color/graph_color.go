package graphcolor

import "fmt"

var maxColor = 3

func graphColor(relations [][]int, color []int, cur int) {
	if cur == len(relations) {
		fmt.Println(color)
		return
	}

	newcolor := make([]int, len(color))
	copy(newcolor, color)
	for i := 1; i <= maxColor; i++ {
		// put color on cur
		newcolor[cur] = i

		// check if color is ok
		if check(relations, newcolor, cur) {
			graphColor(relations, newcolor, cur+1)
		}
	}

}

// check if color is ok
func check(relations [][]int, color []int, cur int) bool {
	for j := range relations[cur] {
		// exists relation between cur and j
		if 1 == relations[cur][j] && color[j] == color[cur] {
			return false
		}
	}
	return true
}

func printRelations(in [][]int) {
	for row := 0; row < len(in); row++ {
		for column := 0; column < len(in[row]); column++ {
			fmt.Printf("%d ", in[row][column])
		}
		fmt.Println()
	}
}
