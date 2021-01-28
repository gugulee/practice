package md

import "math"

var minDist = math.MaxInt16

func minimumDistance(data [][]int, i, j, dist int) {
	length := len(data)

	if i == length || j == length {
		if i == length-1 && j == length {
			if dist < minDist {
				minDist = dist
			}
		}
		return
	}

	if i < length {
		minimumDistance(data, i+1, j, dist+data[i][j])
	}

	if j < length {
		minimumDistance(data, i, j+1, dist+data[i][j])
	}
}
