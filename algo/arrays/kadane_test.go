package arrays

import "testing"

func TestKadane(t *testing.T) {
	original := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	expected := 7
	got := Kadane(original)
	if expected != got {
		t.Fatalf("The Kadane's algorithm implementation has failed returning %v instead of %v", got, expected)
	}
}

func TestKadaneLimits(t *testing.T) {
	original := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	expectedStart := 2
	expectedEnd := 6
	gotStart, gotEnd := KadaneLimits(original)
	if expectedStart != gotStart || expectedEnd != gotEnd {
		t.Fatalf("The Kadane's algorithm implementation has failed returning limits (%v, %v) instead of (%v,%v)",
			gotStart,
			gotEnd,
			expectedStart,
			expectedEnd,
		)
	}
}
