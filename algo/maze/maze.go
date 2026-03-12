package maze

import (
	"fmt"

	set "github.com/deckarep/golang-set"
)

// maze implement by recursion
func maze(maze [][]int, s, t int) {
	length := len(maze)
	// store the prev node,
	// e.g., if prev[3]=1, that is to say, the previous node of 3 is 1
	prev := make([]int, length)
	// store the the node which has been visited
	visited := set.NewSet()

	for i := range prev {
		prev[i] = -1
	}

	if found := recurDfs(maze, prev, visited, s, t); found {
		printPrev(prev, s, t)
	}
}

func recurDfs(maze [][]int, prev []int, visited set.Set, s, t int) bool {
	// add s into visited
	visited.Add(s)

	// if found, print the path and return
	if s == t {
		return true
	}

	for i := range maze[s] {
		if 0 != maze[s][i] && !visited.Contains(i) {
			prev[i] = s
			return recurDfs(maze, prev, visited, i, t)
		}
	}

	return false
}

// maze1 implement by stack
func maze1(maze [][]int, s, t int) {
	length := len(maze)
	// store the prev node,
	// e.g., if prev[3]=1, that is to say, the previous node of 3 is 1
	prev := make([]int, length)
	// store the the node which has been visited
	visited := set.NewSet()

	// stack
	stack := []int{s}

	found := false

	for i := range prev {
		prev[i] = -1
	}

	for 0 != len(stack) {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if node == t {
			found = true
			break
		}

		visited.Add(node)

		for i := range maze[node] {
			if 0 != maze[node][i] && !visited.Contains(i) {
				prev[i] = node
				stack = append(stack, i)
			}
		}
	}

	if found {
		printPrev(prev, s, t)
	}
}

// printPrev print the path s -> t
func printPrev(prev []int, s, t int) {
	if -1 != prev[t] && s != t {
		printPrev(prev, s, prev[t])
	}

	fmt.Printf("%d ", t)
}
