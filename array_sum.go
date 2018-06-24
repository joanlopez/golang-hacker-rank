package main

func ArraySum(ar []int32) int32 {
	var sum int32
	for _, val := range ar {
		sum += val
	}
	return sum
}
