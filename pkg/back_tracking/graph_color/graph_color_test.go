package graphcolor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateTestData() (relation [][]int) {
	relation = make([][]int, 5)
	for i := 0; i < 5; i++ {
		relation[i] = make([]int, 5)
	}

	// vertex A
	relation[0][1] = 1
	relation[0][2] = 1

	// vertex B
	relation[1][0] = 1
	relation[1][2] = 1
	relation[1][3] = 1
	relation[1][4] = 1

	// vertex C
	relation[2][0] = 1
	relation[2][1] = 1
	relation[2][4] = 1

	// vertex D
	relation[3][1] = 1
	relation[3][4] = 1

	// vertex E
	relation[4][1] = 1
	relation[4][2] = 1
	relation[4][3] = 1

	return relation
}

func Test_check(t *testing.T) {
	relations := generateTestData()
	color := make([]int, 5)
	printRelations(relations)
	color[0] = 1
	out := check(relations, color, 0)
	require.True(t, out)

	color[1] = 1
	out = check(relations, color, 0)
	require.False(t, out)

	color[1] = 2
	color[2] = 2
	out = check(relations, color, 0)
	require.True(t, out)

	out = check(relations, color, 1)
	require.False(t, out)

	color[2] = 3
	out = check(relations, color, 0)
	require.True(t, out)
	out = check(relations, color, 1)
	require.True(t, out)

	t.Run("the relation is ok", func(t *testing.T) {
		relations := generateTestData()
		color := make([]int, 5)
		
		color[0]=1
		color[1]=2
		color[2]=3
		color[3]=3
		color[4]=1
	
		out = check(relations, color, 0)
		require.True(t, out)
	
		out = check(relations, color, 1)
		require.True(t, out)
	
		out = check(relations, color, 2)
		require.True(t, out)
	
		out = check(relations, color, 3)
		require.True(t, out)
	
		out = check(relations, color, 4)
		require.True(t, out)
	})

}

func Test_graphColor(t *testing.T) {
	relations := generateTestData()
	color := make([]int, 5)
	printRelations(relations)
	fmt.Println("--------------------------------")
	// put color from 0
	graphColor(relations, color, 0)
}
