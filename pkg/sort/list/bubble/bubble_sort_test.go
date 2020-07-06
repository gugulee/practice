package bubble

import (
	"testing"

	"github.com/stretchr/testify/require"
	singlelist "github.com/practice/pkg/list/single"
)

func TestBubbleSort(t *testing.T) {
	list := singlelist.NewLinkList()
	list.InsertTail(6)
	list.InsertTail(4)
	list.InsertTail(1)
	list.InsertTail(5)
	list.InsertTail(2)
	list.InsertTail(7)
	list.InsertTail(3)

	bubbleSort(list)
	require.Equal(t, "1->2->3->4->5->6->7", list.Print())
}
