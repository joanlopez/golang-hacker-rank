package linear

import "github.com/joanlopez/golang-hacker-rank/strings"

type Queue interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	Empty() bool
	Length() int
	String() string
}

func NewQueue(strFn strings.Fn) Queue {
	return NewDoublyLinkedList(nil, strFn)
}