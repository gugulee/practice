package kahn

import "fmt"

func topoSortByKahn(data [][]bool) {
	length := len(data)
	degree := make([]int, length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if data[i][j] {
				degree[j]++
			}
		}
	}

	fmt.Println(degree)

	var queue []int

	for i := 0; i < length; i++ {
		if 0 == degree[i] {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print("->", cur)

		for i := range data[cur] {
			if !data[cur][i] {
				continue
			}

			degree[i]--
			if 0 == degree[i] {
				queue = append(queue, i)
			}
		}
	}
	fmt.Println()
}

func topoSortByDFS(data [][]bool) {
	length := len(data)
	inverse := make([][]bool, length)
	for i := 0; i < length; i++ {
		inverse[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if data[i][j] {
				inverse[j][i] = true
			}
		}
	}

	visited := make([]bool, length)
	for i := 0; i < length; i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		dfs(inverse, visited, i)
	}
	fmt.Println()
}

func dfs(inverse [][]bool, visited []bool, cur int) {
	for i := range inverse[cur] {
		if visited[i] {
			continue
		}

		if !inverse[cur][i] {
			continue
		}

		visited[i] = true
		dfs(inverse, visited, i)
	}
	fmt.Print("->", cur)
}
