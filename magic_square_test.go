package main

import "testing"

func TestMagicSquare(t *testing.T) {
	origin := [][]int32{
		{4, 8, 2},
		{4, 5, 7},
		{6, 1, 6},
	}
	var want int32 = 4
	got := MagicSquare(origin)
	if want != got {
		t.Errorf("Incorrect magic square min cost, got: %v, want: %v.", got, want)
	}
}
