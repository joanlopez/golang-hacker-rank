package arrays

import "golang-hacker-rank/compare"

func QuickSort(arr []interface{}, low, high int, cmpFn compare.Fn) {
	if low < high {
		pivot := partition(arr, low, high, cmpFn)

		QuickSort(arr, low, pivot-1,  cmpFn)
		QuickSort(arr, pivot+1, high, cmpFn)
	}
}

func partition(arr []interface{}, low, high int, cmpFn compare.Fn) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j < high; j++ {
		if 0 <= cmpFn(pivot, arr[j]) {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
