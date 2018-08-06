package arrays

import "testing"

func TestLeastFrequent(t *testing.T) {
	original := []int{1, 3, 2, 1, 2, 2, 3, 1}
	expected := 3
	got := LeastFrequent(original, 8)
	if expected != got {
		t.Fatalf("The least frequent has failed returning %v instead of %v", got, expected)
	}
}