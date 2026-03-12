package dijkstra

import (
	"fmt"
	"math"
)

func dijkstra(node [][]int, s, t int) {
	lenNode := len(node)
	mw := make([]int, lenNode)
	for i := 0; i < lenNode; i++ {
		if i == s {
			mw[s] = 0
			continue
		}
		mw[i] = math.MaxInt32
	}

	inqueue := make([]bool, lenNode)
	prev := make([]int, lenNode)
	queue := NewPriorityQueue(lenNode)
	queue.add(s, 0)
	inqueue[s] = true
	for !queue.IsEmpty() {
		cur := queue.pull()
		for i := range node[cur.id] {
			if 0 == node[cur.id][i] || i == cur.id {
				continue
			}
			tmp := mw[cur.id] + node[cur.id][i]
			if tmp < mw[i] {
				mw[i] = tmp
				prev[i] = cur.id
				if inqueue[i] {
					queue.update(i, mw[i])
				} else {
					queue.add(i, mw[i])
					inqueue[i] = true
				}
			}
		}
	}

	fmt.Println(mw)
	fmt.Printf("the path from %d to %d is: ", s, t)
	printPath(prev, s, t)
	fmt.Println()
	fmt.Println("the shortest path is:", mw[t])
}

func printPath(prev []int, s, t int) {
	if s == t {
		fmt.Print(s)
		return
	}
	printPath(prev, s, prev[t])
	fmt.Print("->", t)
}
