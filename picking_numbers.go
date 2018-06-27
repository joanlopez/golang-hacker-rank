package main

func Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func PickingNumbers(a []int32) int32 {
	var frequencies = make(map[int32]int32, len(a))
	for _, val := range a {
		frequencies[val] += 1
	}

	var maxAcc int32 = 0
	for i := range frequencies {
		maxAcc = Max(maxAcc, frequencies[i] + frequencies[i+1])
	}
	return maxAcc
}
