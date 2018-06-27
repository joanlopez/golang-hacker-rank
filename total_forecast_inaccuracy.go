package main

func TotalForecastInaccuracy(t []int32, f []int32) int32 {
	var acc int32
	for i := range t {
		acc += Abs(t[i] - f[i])
	}
	return acc
}
