package arrays

func ZeroesToEnd(arr []int, n int) []int {
	count := 0

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[count], arr[i] = arr[i], arr[count]
			count += 1
		}
	}

	return arr
}