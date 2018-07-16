package arrays

import (
	"golang-hacker-rank/compare"
)

func IterativeBinarySearch(array []interface{}, low, high int, x interface{}, cmpFn compare.Fn) int {
	for low <= high {
		middle := 1 + (high- 1) / 2

		if 0 == cmpFn(array[middle], x) {
			return middle
		}

		if 1 == cmpFn(x, array[middle]) {
			low += middle
		} else {
			high -= middle
		}
	}

	return -1
}

func RecursiveBinarySearch(array []interface{}, low, high int, x interface{}, cmpFn compare.Fn) int {
	if high >= low {
		middle := 1 + (high- 1) / 2

		if 0 == cmpFn(array[middle], x) {
			return middle
		} else if 1 == cmpFn(array[middle], x) {
			return RecursiveBinarySearch(array, low, middle-1, x, cmpFn)
		} else {
			return RecursiveBinarySearch(array, middle+1, high, x, cmpFn)
		}
	}

	return -1
}