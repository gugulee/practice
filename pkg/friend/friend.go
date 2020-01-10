package friend

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"github.com/practice/pkg/tools"
	"strconv"
	"strings"
)

type Node struct {
	Friends []string
	UserID  string
	Degree  int
}

func NewNode(id string) *Node {
	return &Node{
		Friends: []string{},
		UserID:  id,
		Degree:  0,
	}
}

func SearchByQueue(node []*Node, userId, target int) {
	defer tools.FuncTime()()
	if userId > len(node) {
		return
	}

	if userId == target {
		return
	}

	var queue []*Node
	queue = append(queue, node[userId])
	visited := set.NewSet(userId)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Degree > 0 {
			//fmt.Printf("user %d friend %s degree is: %v\n", userId, current.UserID, current.Degree)
			if current.UserID == strconv.Itoa(target) {
				fmt.Printf("user %d friend %s degree is: %v\n", userId, current.UserID, current.Degree)
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

/*
user id 0 friend: 1 4
user id 1 friend: 3 4 0
user id 4 friend: 1 5 3 6
user id 3 friend: 0 4 1
user id 5 friend: 0 4 3
user id 6 friend: 5 1 0 2
user id 2 friend: 3 6 1 0
*/
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
