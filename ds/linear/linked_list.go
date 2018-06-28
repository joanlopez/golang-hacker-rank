package linear

import "reflect"

type ListNode struct {
	Value interface{}
	Next *ListNode
}

type List interface {
	Add()
	Len() int
	Contains(value interface{}) bool
	GetAt(at int) interface{}
}

func CreateListNode(value interface{}) *ListNode {
	return &ListNode{
		value,
		nil,
	}
}

func (head *ListNode) Add(post *ListNode) {
	// TODO: Define what happens when head == nil
	for head.Next != nil {
		head = head.Next
	}
	head.Next = post
}

func (head *ListNode) Len() int {
	if head == nil {
		return 0
	}

	sum := 1
	for head.Next != nil {
		sum++
		head = head.Next
	}
	return sum
}

// TODO: Review the time complexity of DeepEquals operation
func (head *ListNode) Contains(value interface{}) bool {
	contains := false
	for head != nil {
		contains = reflect.DeepEqual(head.Value, value)
		if contains {
			break
		}
		head = head.Next

	}
	return contains
}

// TODO: Review the time complexity of DeepEquals operation
func (head *ListNode) GetAt(at int) interface{} {
	var obtained interface{}
	pos := 0
	for head != nil {
		if at == pos {
			obtained = head.Value
			break
		}
		head = head.Next
		pos++
	}
	return obtained
}