package linear

import (
	"testing"
	"github.com/joanlopez/golang-hacker-rank/compare"
)

func TestStack(t *testing.T) {
	stack := buildStack()
	stack.Stack(4)
	if 4 != stack.Length() {
		t.Errorf("Error on Stack() method")
	}
}

func TestUnstack(t *testing.T) {
	errMsg := "Error on Unstack() method"
	stack := buildStack()
	if v := stack.Unstack(); 0 != compare.IntFn(3, v) {
		t.Errorf(errMsg)
	}
	if v := stack.Unstack(); 0 != compare.IntFn(2, v) {
		t.Errorf(errMsg)
	}
	if v := stack.Unstack(); 0 != compare.IntFn(1, v) {
		t.Errorf(errMsg)
	}
	if true != stack.Empty() {
		t.Errorf(errMsg)
	}
}

// Builds the following stack:
//
// 3
// |
// 2
// |
// 1
//
func buildStack() Stack {
	stack := NewStack(nil)
	stack.Stack(1)
	stack.Stack(2)
	stack.Stack(3)
	return stack
}
