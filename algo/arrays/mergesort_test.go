package arrays

import (
	"golang-hacker-rank/compare"
	"testing"
	"reflect"
)

type MergeSortTestCase struct {
	name     string
	original []interface{}
	expected []interface{}
	cmpFn    compare.Fn
}

var mergeSortTestCases = []MergeSortTestCase{
	{"array of ints", []interface{}{3, 5, 4, 1, 2}, []interface{}{1, 2, 3, 4, 5}, compare.IntFn},
	{"array of strings", []interface{}{"c", "e", "d", "a", "b"}, []interface{}{"a", "b", "c", "d", "e"}, compare.StringFn},
}

func TestMergeSort(t *testing.T) {
	for _, tc := range mergeSortTestCases {
		t.Run(tc.name, func(t *testing.T) {
			MergeSort(tc.original, 0, 4, tc.cmpFn)
			if !reflect.DeepEqual(tc.expected, tc.original) {
				t.Fatalf("mergesort %v has failed returning %v instead of %v", tc.name, tc.original, tc.expected)
			}
		})
	}

}
