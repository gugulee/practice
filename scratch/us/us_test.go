package us

import (
	"testing"
)

type LargeItem struct {
	id  int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]LargeItem
	for b.Loop() {
		tmp := 0
		for j := 0; j < len(items); j++ {
		// for j := range len(items) {
			tmp = items[j].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]LargeItem
	for b.Loop() {
		tmp := 0
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
