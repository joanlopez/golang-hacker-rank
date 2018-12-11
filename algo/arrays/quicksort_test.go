package arrays

import (
	"github.com/joanlopez/golang-hacker-rank/compare"
	"reflect"
	"testing"
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
	for _, tc := range quickSortTestCases {
		t.Run(tc.name, func(t *testing.T) {
			QuickSort(tc.original, 0, 4, tc.cmpFn)
			if !reflect.DeepEqual(tc.expected, tc.original) {
				t.Fatalf("quicksort %v has failed returning %v instead of %v", tc.name, tc.original, tc.expected)
			}
		})
	}

}
