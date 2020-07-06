package bubble

import (

	singlelist "github.com/practice/pkg/list/single"
)

func bubbleSort(original *singlelist.LinkList) {
	// except for less than 2 node
	if nil == original.Head() || nil == original.Head().Next() || nil == original.Head().Next().Next() {
		return
	}
	for i := 0; i < int(original.Length()); i++ {
		flag := false
		prev := original.Head()
		node := original.Head().Next()
		for nil != node.Next() {
			nodeValue := node.Value().(int)
			nextValue := node.Next().Value().(int)
			if nodeValue > nextValue {
				tmp := node.Next()
				original.Swap(prev, node)
				flag = true
				node = tmp
			}
			prev = node
			node = node.Next()
		}
		if !flag {
			break
		}
	}
}
