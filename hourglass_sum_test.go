package main

import (
	"testing"
)

func TestHourglassSum(t *testing.T) {
	origin := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	var want int32 = 19
	got := HourglassSum(origin)
	if want != got {
		t.Errorf("Incorrect hourglass sum, got: %v, want: %v.", got, want)
	}
}
