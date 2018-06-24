package main

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func DiagonalDifference(arr [][]int32) int32 {
	var d1sum int32
	var d2sum int32
	var arrLen = len(arr)
	for i := 0; i < arrLen; i++ {
		d1sum += arr[i][i]
		d2sum += arr[i][arrLen-i-1]
	}
	return abs(d1sum - d2sum)
}
