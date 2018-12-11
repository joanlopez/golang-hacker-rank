package arrays

import (
	"github.com/joanlopez/golang-hacker-rank/strings"
	"testing"
)

func TestQuadraticSlidingWindowMax(t *testing.T) {
	original := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	k := 3
	expected := "3 3 4 5 5 5 6"
	got := QuadraticSlidingWindowMax(original, 9, k).String(strings.IntFn)
	if expected != got {
		t.Fatalf("The quadratic sliding window max has failed returning %v instead of %v", got, expected)
	}
}
