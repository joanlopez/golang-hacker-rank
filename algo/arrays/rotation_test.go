package arrays

import (
	"testing"
	"reflect"
)

type RotateFn func(arr []int, d, n int)

type RotationTestCase struct {
	name     string
	original []int
	moves    int
	length   int
	fn       RotateFn
	expected []int
}

// TODO: If the cases will always be the same, then it could consist a simple array of functions
var rotateTestCases = []RotationTestCase{
	{"reverse", []int{1, 2, 3, 4, 5, 6, 7}, 2, 7, ReverseRotation, []int{3, 4, 5, 6, 7, 1, 2}},
	{"one by one", []int{1, 2, 3, 4, 5, 6, 7}, 2, 7, OneByOneRotation, []int{3, 4, 5, 6, 7, 1, 2}},
	{"temporary", []int{1, 2, 3, 4, 5, 6, 7}, 2, 7, TemporaryRotation, []int{3, 4, 5, 6, 7, 1, 2}},
	{"juggling", []int{1, 2, 3, 4, 5, 6, 7}, 2, 7, JugglingRotation, []int{3, 4, 5, 6, 7, 1, 2}},
}

func TestRotate(t *testing.T) {
	for _, tc := range rotateTestCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.fn(tc.original, tc.moves, tc.length)
			if !reflect.DeepEqual(tc.expected, tc.original) {
				t.Fatalf("%v rotation has failed returning %v instead of %v",
					tc.name, tc.original, tc.expected)
			}
		})
	}
}
