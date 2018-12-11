package arrays

import (
	"github.com/joanlopez/golang-hacker-rank/compare"
	"github.com/joanlopez/golang-hacker-rank/ds/linear"
	"github.com/joanlopez/golang-hacker-rank/strings"
)

func QuadraticSlidingWindowMax(arr []int, arraySize, intervalSize int) *linear.DoublyLinkedList {
	var max int
	maximums := linear.NewDoublyLinkedList(compare.IntFn, strings.IntFn)

	for i := 0; i <= arraySize-intervalSize; i++ {
		max = arr[i]

		for j := 1; j < intervalSize; j++ {
			if arr[i+j] > max {
				max = arr[i+j]
			}
		}
		maximums.Append(max)
	}
	return maximums
}
