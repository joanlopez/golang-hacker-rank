package arrays

func LeastFrequent(arr []int, n int) int {
	count := make(map[int]int, n)

	for _, x := range arr {
		if _, ok := count[x]; !ok {
			count[x] = 0
		}
		count[x] += 1
	}

	min := n+1; res := -1

	for key, value := range count {
		if value < min {
			min = value
			res = key
		}
	}

	return res
}