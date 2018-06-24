package main

// O(n) or O(n/2 ?
func ReverseArray(arr []int32) []int32 {
	length := len(arr)
	length2 := length >> 1
	var temp int32

	for i := 0; i < length2; i++ {
		temp = arr[i]
		arr[i] = arr[length- i - 1]
		arr[length- i - 1] = temp
	}

	return arr
}

