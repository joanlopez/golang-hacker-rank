package arrays

import (
	"golang-hacker-rank/compare"
	"testing"
	"reflect"
)

type QuickSortTestCase struct {
	name     string
	original []interface{}
	expected []interface{}
	cmpFn    compare.Fn
}

var quickSortTestCases = []QuickSortTestCase{
	{"array of ints", []interface{}{3, 5, 4, 1, 2}, []interface{}{1, 2, 3, 4, 5}, compare.IntFn},
	{"array of strings", []interface{}{"c", "e", "d", "a", "b"}, []interface{}{"a", "b", "c", "d", "e"}, compare.StringFn},
}

func TestQuickSort(t *testing.T) {
	original := []interface{}{3, 5, 4, 1, 2}
	expected := []interface{}{1, 2, 3, 4, 5}
	QuickSort(original, 0, 4, compare.IntFn)

	for _, tc := range quickSortTestCases {
		t.Run(tc.name, func(t *testing.T) {
			QuickSort(tc.original, 0, 4, tc.cmpFn)
			if !reflect.DeepEqual(tc.expected, tc.original) {
				t.Fatalf("quicksort %v has failed returning %v instead of %v", tc.name, original, expected)
			}
		})
	}

}
