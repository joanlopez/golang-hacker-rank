package linear

import "github.com/joanlopez/golang-hacker-rank/strings"

type Stack interface {
	Stack(data interface{})
	Unstack() interface{}
	Empty() bool
	Length() int
	String() string
}

func NewStack(strFn strings.Fn) Stack {
	return NewDoublyLinkedList(nil, strFn)
}