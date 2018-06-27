package main

import "testing"

func TestPickingNumbers(t *testing.T) {
	origin := []int32{1, 2, 2, 3, 1, 2}
	var want int32 = 5
	got := PickingNumbers(origin)
	if want != got {
		t.Errorf("Incorrect picking numbers result, got: %v, want: %v.", got, want)
	}
}
