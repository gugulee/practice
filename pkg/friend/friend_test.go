package friend

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

func TestSearchByQueue(t *testing.T) {
	usersNode := generateRelation(100, 200)
	//for _, user := range usersNode {
	//	fmt.Printf("user %s friend is %s\n", user.UserID, user.Friends)
	//}

	bfs(usersNode, 0, 2)
}

func TestPrint(t *testing.T) {
	usersNode := generateRelation(7, 11)
	for _, user := range usersNode {
		fmt.Printf("user %s friend is %s\n", user.UserID, user.Friends)
	}
	fmt.Println("-------------------")
	Print(usersNode, 0)
}
