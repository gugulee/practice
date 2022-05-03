package sharesweets

import (
	"fmt"
	"testing"
)

func Test_shareSweets(t *testing.T) {
	sweetsSize := []int{3, 5, 7, 10, 13}
	childrenNeed := []int{2, 5, 6, 8, 11, 12, 13}
	shareSweets(sweetsSize, childrenNeed)
	fmt.Println("------------------------")
	sweetsSize = []int{3, 5, 7, 10, 13}
	childrenNeed = []int{2, 5, 6, 8, 14, 15, 20}
	shareSweets(sweetsSize, childrenNeed)
}
