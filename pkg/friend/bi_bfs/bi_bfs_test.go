package bi_bfs

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func addFriend(all *[]string, id int) {
	for i := range *all {
		if strconv.Itoa(id) == (*all)[i] {
			return
		}
	}

	*all = append(*all, strconv.Itoa(id))
}

func generateRelation(userMax int, relations int) []*Node {
	var usersNode = make([]*Node, userMax)

	for i := 0; i < userMax; i++ {
		usersNode[i] = NewNode(strconv.Itoa(i))
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i <= relations; i++ {
		aId := rand.Intn(userMax)
		bId := rand.Intn(userMax)
		if aId == bId {
			continue
		}

		aUser := usersNode[aId]
		bUser := usersNode[bId]

		addFriend(&aUser.Friends, bId)
		addFriend(&bUser.Friends, aId)

	}
	return usersNode
}

func generateRelation1(userMax int, relations map[int][]string) []*Node {
	var usersNode = make([]*Node, userMax)

	for i := 0; i < userMax; i++ {
		usersNode[i] = NewNode(strconv.Itoa(i))
	}

	for id, friends := range relations {
		usersNode[id].Friends = friends
	}

	return usersNode
}

func TestBidirectionalBfs(t *testing.T) {
	//relations := map[int][]string{
	//	0: {"1"},
	//	1: {"0","13"},
	//	2: {"3"},
	//	3: {"2", "4"},
	//	4: {"3", "5"},
	//	5: {"4", "6"},
	//	6: {"5", "7"},
	//	7: {"6", "8"},
	//	8: {"7", "9"},
	//	9: {"8", "10"},
	//	10: {"9", "11"},
	//	11: {"10", "12"},
	//	12: {"11", "13"},
	//	13: {"12", "1"},
	//}
	//
	//usersNode := generateRelation1(14, relations)
	usersNode := generateRelation(30000000, 50000000)

	//for _, user := range usersNode {
	//	fmt.Printf("user %s friend is %s\n", user.UserID, user.Friends)
	//}

	s, d := 0, 2
	out := bidirectionalBfs(s, d, usersNode)
	fmt.Printf("user %d friend %d degree is: %d\n", s, d, out)
}
