package implementation

import (
	"testing"
)

// TODO: Improve the testing suite (table test approach)
func TestCountFruit(t *testing.T) {
	var start int32 = 7
	var end int32 = 11

	var aTree int32 = 5
	var oTree int32 = 15

	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	if count := CountFruit(start, end, aTree, apples); count != 1 {
		t.Errorf("Incorrect fruit count, got: %v, want: %v.", count, 1)
	}

	if count := CountFruit(start, end, oTree, oranges); count != 1 {
		t.Errorf("Incorrect fruit count, got: %v, want: %v.", count, 1)
	}
}

func TestCountApplesAndOranges(t *testing.T) {
	CountApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
