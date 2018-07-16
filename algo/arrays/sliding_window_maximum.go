package arrays

import "golang-hacker-rank/ds/linear"

func QuadraticSlidingWindowMax(arr []int, arraySize, intervalSize int) *linear.LinkedList {
	var max int
	maximums := &linear.LinkedList{}

	for i := 0; i <= arraySize-intervalSize; i++ {
		max = arr[i]

		for j := 1; j < intervalSize; j++ {
			if arr[i+j] > max {
				max = arr[i+j]
			}
		}
		maximums.InsertLast(max)
	}
	return maximums
}
