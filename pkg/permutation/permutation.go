package permutation

import "fmt"

// in=[]int{1,2,3}, idx in the max process index
func permutation(in []int, idx int) {
	if 0 == idx {
		fmt.Println(in)
		return
	}

	for i := 0; i <= idx; i++ {
		in[i], in[idx] = in[idx], in[i]

		permutation(in, idx-1)

		in[i], in[idx] = in[idx], in[i]
	}
}
