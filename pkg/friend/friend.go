package friend

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"github.com/practice/pkg/tools"
	"strconv"
	"strings"
)

// Node ...
type Node struct {
	Friends []string
	UserID  string
	Degree  int
}

// NewNode ...
func NewNode(id string) *Node {
	return &Node{
		Friends: []string{},
		UserID:  id,
		Degree:  0,
	}
}

// bfs ...
func bfs(node []*Node, userID, target int) {
	defer tools.FuncTime()()
	if userID > len(node) {
		return
	}

	if userID == target {
		return
	}

	var queue []*Node
	queue = append(queue, node[userID])
	visited := set.NewSet(userID)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Degree > 0 {
			//fmt.Printf("user %d friend %s degree is: %v\n", userId, current.UserID, current.Degree)
			if current.UserID == strconv.Itoa(target) {
				fmt.Printf("user %d friend %s degree is: %v\n", userID, current.UserID, current.Degree)
				break
			}
		}

		for _, friend := range current.Friends {
			id, _ := strconv.Atoi(friend)
			if visited.Contains(id) {
				continue
			}
			node[id].Degree = current.Degree + 1
			queue = append(queue, node[id])
			visited.Add(id)
		}
	}
}

// Print ...
func Print(node []*Node, id int) {
	if id > len(node) {
		return
	}

	var queue []*Node
	queue = append(queue, node[id])
	visited := set.NewSet()

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited.Contains(current.UserID) {
			continue
		}

		fmt.Printf("user id %s friend: %v\n", current.UserID, printSlice(current.Friends))

		for _, friend := range current.Friends {
			id, _ := strconv.Atoi(friend)
			queue = append(queue, node[id])
		}
		visited.Add(current.UserID)
	}
}

func printSlice(s []string) string {
	return strings.Join(s, " ")
}
