package main

import (
	"math"
)

// Fixed solution for 3x3 squares
// The key point is taking into account the unique magic number
// and the possible solutions (permutations) for it:

var posPerms = [][][]int32{
	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
}

func Abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func Min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func MagicSquare(s [][]int32) int32 {
	var minCost int32 = math.MaxInt32
	for perm := 0; perm < 8; perm++ {
		var permCost int32 = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				permCost += Abs(s[i][j] - posPerms[perm][i][j])
			}
		}
		minCost = Min(minCost, permCost)
	}
	return minCost
}
