package arrays

import (
	"testing"
	"golang-hacker-rank/compare"
)

type tc struct {
	name     string
	arr      []interface{}
	search   interface{}
	expected int
	fn       func(array []interface{}, low, high int, x interface{}, cmpFn compare.Fn) int
	cmpFn    compare.Fn
}

var tests = []tc{
	{"iterative existing int", []interface{}{1, 2, 3, 4, 5}, 3, 2, IterativeBinarySearch, compare.IntFn},
	{"iterative non-existing int", []interface{}{1, 2, 4, 5, 6}, 3, -1, IterativeBinarySearch, compare.IntFn},
	{"iterative existing string", []interface{}{"a", "b", "c", "d", "e"}, "c", 2, IterativeBinarySearch, compare.StringFn},
	{"iterative non-existing string", []interface{}{"a", "b", "d", "e", "f"}, "c", -1, IterativeBinarySearch, compare.StringFn},
	{"recursive existing int", []interface{}{1, 2, 3, 4, 5}, 3, 2, RecursiveBinarySearch, compare.IntFn},
	{"recursive non-existing int", []interface{}{1, 2, 4, 5, 6}, 3, -1, RecursiveBinarySearch, compare.IntFn},
	{"recursive existing string", []interface{}{"a", "b", "c", "d", "e"}, "c", 2, RecursiveBinarySearch, compare.StringFn},
	{"recursive non-existing string", []interface{}{"a", "b", "d", "e", "f"}, "c", -1, RecursiveBinarySearch, compare.StringFn},
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.fn(tc.arr, 0, 4, tc.search, tc.cmpFn)
			if tc.expected != got {
				t.Fatalf("binary search %v with %v has failed returning %v instead of %v",
					tc.name, tc.arr, got, tc.expected)
			}
		})
	}
}
