package arrays

import "github.com/joanlopez/golang-hacker-rank/math"

func Kadane(arr []int) (globalMax int) {
	// We assume that len(slice) is O(1), otherwise argument should be given
	size := len(arr)
	globalMax = arr[0]; localMax := arr[0]
	for i := 1; i < size; i++ {
		localMax = math.Max(arr[i], localMax+arr[i])
		globalMax = math.Max(localMax, globalMax)
	}
	return
}

func KadaneLimits(arr []int) (start, end int) {
	// We assume that len(slice) is O(1), otherwise argument should be given
	size := len(arr)
	start = 0; end = 0; s := 0
	globalMax := math.MinInt; localMax := 0
	for i := 0; i < size; i++ {
		localMax += arr[i]
		if localMax < 0 {
			localMax = 0
			s = i + 1
		} else if localMax > globalMax {
			globalMax = localMax
			start = s
			end = i
		}
	}
	return
}
