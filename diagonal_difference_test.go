package main

import "testing"

func TestDiagonalDifference(t *testing.T) {
	origin := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	var want int32 = 15
	got := DiagonalDifference(origin)
	if want != got {
		t.Errorf("Incorrect diagonal difference, got: %v, want: %v.", got, want)
	}
}