package main

import (
	"sort"
	"fmt"
)

func sumTo64(arr []int32) int64 {
	var sum int64
	for _, val := range arr {
		sum += int64(val)
	}
	return sum
}
// Analyze the possible sort algorithms
func MiniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arrLen := len(arr)
	min := sumTo64(arr[:4])
	max := sumTo64(arr[arrLen-4:])
	fmt.Printf("%v %v", min, max)
}
