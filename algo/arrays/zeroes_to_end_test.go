package arrays

import (
	"reflect"
	"testing"
)

func TestZeroesToEnd(t *testing.T) {
	original := []int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9}
	expected := []int{1, 9, 8, 4, 2, 7, 6, 9, 0, 0, 0, 0}
	got := ZeroesToEnd(original, 12)

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("zeroes to end has failed returning %v instead of %v", got, expected)
	}
}
