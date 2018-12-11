package implementation

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	original := []int32 {73, 67, 38, 33}
	want := []int32 {75, 67, 40, 33}
	got := GradingStudents(original)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Incorrect grading students, got: %v, want: %v.", got, want)
	}
}
