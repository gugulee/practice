package permutation

import "testing"

func TestPermutate(t *testing.T) {
	in := []int{1, 2, 3}
	permutation(in, len(in)-1)
}
