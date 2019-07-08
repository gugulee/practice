package single

import (
	"fmt"
	"github.com/practice/pkg/list/single"
)

type LruSingle struct {
	length uint32
	lru    *single.LinkList
}

func NewLru(length uint32) *LruSingle {
	return &LruSingle{length: length, lru: single.NewLinkList()}
}

func (lru *LruSingle) Length() uint32 {
	return lru.length
}

func (lru *LruSingle) Lru() *single.LinkList {
	return lru.lru
}

/*
here's my way of thinking about lru: maintain an ordered single list(the list is fixed-size),
we traverse the list sequentially from the list header,When new data come.
There are some conditions,as follows:
1.If the data does not exist in the list, there are two cases:
(1)If the list is not full, insert the new data into the list header
(2)If the list is full, delete the tail node and insert the new data into the list header
2.If the data exists in the list, delete the node correspond with the data and insert into the list header
*/
func (lru *LruSingle) LruCache(value string) (bool, error) {
	list := lru.lru
	if nil == list {
		return false, fmt.Errorf("the head of list can no be empty")
	}

	if "" == value {
		return false, fmt.Errorf("the value can no be empty")
	}

	node := list.SearchListByValue(value)
	// node is nil means that not find value
	if nil == node {
		if list.Length() == lru.length {
			node = list.Head()
			for ; nil != node.Next; node = node.Next {

			}
		}
	}

	if nil != node && ! list.DeleteNode(node) {
		return false, fmt.Errorf("delete the tail:%s failed", node.Value().(string))
	}

	list.InsertHead(value)
	return true, nil
}

func (lru *LruSingle) Print() string {
	return lru.Lru().Print()
}
