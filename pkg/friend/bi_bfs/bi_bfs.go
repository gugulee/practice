package bi_bfs

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"github.com/practice/pkg/tools"
	"strconv"
)

type Node struct {
	Friends []string
	UserID  string

	// the degree that other node to this node
	Degrees map[string]int
}

func NewNode(id string) *Node {
	return &Node{
		Friends: []string{},
		UserID:  id,
		Degrees: map[string]int{id: 0},
	}
}

func getNextDegreeFriends(allNodes, queue []*Node, visited set.Set, degree *int) []*Node {
	var newQueue []*Node
	for 0 != len(queue) {
		current := queue[0]
		queue = queue[1:]

		for _, friend := range current.Friends {
			id, _ := strconv.Atoi(friend)
			if visited.Contains(id) {
				continue
			}

			visited.Add(id)
			newQueue = append(newQueue, allNodes[id])
		}
	}

	*degree++
	return newQueue
}

func bidirectionalBfs(a, b int, nodes []*Node) int {
	defer tools.FuncTime()()
	if a == b {
		return 0
	}

	// the max degreee of a and b, which can not exceed it
	maxDegree := 100
	aDegree, bDegree := 0, 0

	aNodes := []*Node{nodes[a]}
	bNodes := []*Node{nodes[b]}

	aVisited := set.NewSet(a)
	bVisited := set.NewSet(b)
	empty := set.NewSet()

	for aDegree+bDegree < maxDegree {
		if 0 != len(aNodes) {
			aNodes = getNextDegreeFriends(nodes, aNodes, aVisited, &aDegree)
			// if exist intersection, return the sum of degree
			if intersection := aVisited.Intersect(bVisited); !intersection.Equal(empty) {
				fmt.Printf("aDegree=%d,bDegree=%d\n", aDegree, bDegree)
				return aDegree + bDegree
			}
		}

		if 0 != len(bNodes) {
			bNodes = getNextDegreeFriends(nodes, bNodes, bVisited, &bDegree)
			// if exist intersection, return the sum of degree
			if intersection := aVisited.Intersect(bVisited); !intersection.Equal(empty) {
				fmt.Printf("aDegree=%d,bDegree=%d\n", aDegree, bDegree)
				return aDegree + bDegree
			}
		}

		// if aNodes and bNodes are all empty, return -1
		if 0 == len(aNodes) || 0 == len(bNodes) {
			return -1
		}
	}

	return -1
}
