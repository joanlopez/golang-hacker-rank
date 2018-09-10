package linear

import (
	"testing"
	"github.com/joanlopez/golang-hacker-rank/compare"
)

func TestEnqueue(t *testing.T) {
	queue := buildQueue()
	queue.Enqueue(4)
	if 4 != queue.Length() {
		t.Errorf("Error on Enqueue() method")
	}
}

func TestDequeue(t *testing.T) {
	errMsg := "Error on Dequeue() method"
	queue := buildQueue()
	if v := queue.Dequeue(); 0 != compare.IntFn(1, v) {
		t.Errorf(errMsg)
	}
	if v := queue.Dequeue(); 0 != compare.IntFn(2, v) {
		t.Errorf(errMsg)
	}
	if v := queue.Dequeue(); 0 != compare.IntFn(3, v) {
		t.Errorf(errMsg)
	}
	if true != queue.Empty() {
		t.Errorf(errMsg)
	}
}

// Builds the following queue:
//
// 1 - 2 - 3
//
func buildQueue() Queue {
	queue := NewQueue(nil)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	return queue
}
