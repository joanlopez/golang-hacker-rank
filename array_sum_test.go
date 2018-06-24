package main

import "testing"

func TestArraySum(t *testing.T) {
	origin := []int32{1,2,3,4,5}
	var want int32 = 15
	got := ArraySum(origin)
	if want != got {
		t.Errorf("Incorrect array sum, got: %v, want: %v.", got, want)
	}
}
