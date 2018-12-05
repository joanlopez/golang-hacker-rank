package implementation

// TODO: Improve, at least with memoization / dynamic programming
func ViralAdvertising(n int32) int32 {
	var cumulative int32 = 0
	for n > 0 {
		cumulative += liked(n)
		n--
	}
	return cumulative
}

func shared(n int32) int32 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 5
	}

	return liked(n-1) * 3
}

func liked (n int32) int32 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 2
	}

	return shared(n) / 2
}