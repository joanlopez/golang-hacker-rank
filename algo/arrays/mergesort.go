package arrays

import "github.com/joanlopez/golang-hacker-rank/compare"

func MergeSort(arr []interface{}, low, high int, cmpFn compare.Fn) {
	if low < high {
		middle := (low + (high - 1)) / 2

		MergeSort(arr, low, middle, cmpFn)
		MergeSort(arr, middle+1, high, cmpFn)
		merge(arr, low, middle, high, cmpFn)
	}
}

func merge(arr []interface{}, low, middle, high int, cmpFn compare.Fn) {
	n1 := middle - low + 1
	n2 := high - middle

	// Create temporary arrays
	L := make([]interface{}, n1)
	R := make([]interface{}, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[low + i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[middle + 1 + j]
	}

	// Merge the temporary arrays back into arr
	i := 0; j := 0; k := low

	for i < n1 && j < n2 {
		if cmpFn(L[i], R[j]) < 1 {
			arr[k] = L[i]
			i += 1
		} else {
			arr[k] = R[j]
			j += 1
		}
		k += 1
	}

	for i < n1 {
		arr[k] = L[i]
		i += 1
		k += 1
	}

	for j < n2 {
		arr[k] = R[j]
		j += 1
		k += 1
	}
}