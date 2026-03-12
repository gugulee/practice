package matrixmultiply

import "fmt"

//	a=[][]{
//		{1,2},
//	 {3,4},
//	}
//
//	b=[][]{
//		{4,3},
//	 {2,1},
//	}
//
//	c=[][]{
//		{8,5},
//	 {20,13},
//	}
//
// c[0][0]=a[0][0]*b[0][0]+a[0][1]*b[1][0]
// c[0][1]=a[0][0]*b[0][1]+a[0][1]*b[1][1]
// c[1][0]=a[1][0]*b[0][0]+a[1][1]*b[1][0]
// c[1][1]=a[1][1]*b[1][1]+a[1][1]*b[1][1]
func matrixMultiply(a, b [][]int) (c [][]int) {
	c = make([][]int, len(a))
	for i := range c {
		c[i] = make([]int, len(a[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			for k := 0; k < len(a); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func divideAndConquer(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, n)
	}

	if 1 == n {
		c[0][0] = a[0][0] * b[0][0]
		return c
	}

	return c
}

func print(a, b [][]int) {
	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		fmt.Println(a[i], b[j])
	}
}
