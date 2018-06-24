package main

import (
	"math"
)

func sum(params ...int32) int32 {
	return ArraySum(params)
}

// O(X) ?
func HourglassSum(arr [][]int32) int32 {
	var maxSum int32 = math.MinInt32
	maxRow := len(arr) - 2
	maxCol := len(arr[0]) - 2
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			sum := sum(
				arr[row][col], arr[row][col+1], arr[row][col+2],
				arr[row+1][col+1],
				arr[row+2][col], arr[row+2][col+1], arr[row+2][col+2],
			)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
