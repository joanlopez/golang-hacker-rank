package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	original := []int32 {1, 4, 3, 2}
	want := []int32 {2, 3, 4, 1}
	got := ReverseArray(original)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Incorrect reversed array, got: %v, want: %v.", got, want)
	}
}