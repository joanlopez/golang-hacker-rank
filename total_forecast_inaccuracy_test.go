package main

import "testing"

func TestTotalForecastInaccuracy(t *testing.T) {
	originT := []int32{14,13, 12, 13, 16, 18, 21}
	originF := []int32{15, 11, 12, 11, 16, 19, 24}
	var want int32 = 9
	got := TotalForecastInaccuracy(originT, originF)
	if want != got {
		t.Errorf("Incorrect total forecast inaccuracy, got: %v, want: %v.", got, want)
	}
}
