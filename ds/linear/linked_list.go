package linear

import (
	"reflect"
	"errors"
	"bytes"
	"golang-hacker-rank/strings"
)

const emptyListMsg = "trying to do an operation over an empty list"
const emptyValuesMsg = "the array of values MUST contain at least one element"
const nonValidInsertionMsg = "trying to insert a non-valid value"
const nonValidPositionMsg = "the position must be between 1 and Len() (both included)"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func CreateLinkedList(values []interface{}) (*LinkedList, error) {
	valuesLength := len(values)
	if valuesLength == 0 {
		return nil, errors.New(emptyValuesMsg)
	}

	head := &ListNode{Value: values[0]}
	if 1 == valuesLength {
		return &LinkedList{Head: head}, nil
	}

	tail := head
	for _, value := range values[1:] {
		next := &ListNode{Value: value}
		tail.Next = next
		tail = next
	}
	return &LinkedList{head, tail}, nil
}

// Time complexity: O(1)
func (list *LinkedList) Empty() bool {
	return nil == list || nil == list.Head
}

// Time complexity: O(1)
func (list *LinkedList) InsertFirst(value interface{}) (*LinkedList, error) {
	if nil == list {
		return nil, errors.New(emptyListMsg)
	}

	if nil == value {
		return nil, errors.New(nonValidInsertionMsg)
	}

	newNode := &ListNode{Value: value, Next: list.Head}
	list.Head = newNode
	return list, nil
}

// Time complexity: O(1)
func (list *LinkedList) InsertLast(value interface{}) (*LinkedList, error) {
	if nil == list {
		return nil, errors.New(emptyListMsg)
	}

	if nil == value {
		return nil, errors.New(nonValidInsertionMsg)
	}

	newNode := &ListNode{Value: value}

	if nil == list.Head {
		list.Head = newNode
	} else if nil == list.Tail {
		list.Head.Next = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}

	return list, nil
}

// Time complexity: O(n)
func (list *LinkedList) InsertAt(pos int, value interface{}) (*LinkedList, error) {
	if nil == list {
		return nil, errors.New(emptyListMsg)
	}

	if nil == value {
		return nil, errors.New(nonValidInsertionMsg)
	}

	// TODO: Implement it
	currPos := 1
	currElem := list.Head
	for currPos != pos && nil != currElem.Next {
		currPos++
		currElem = currElem.Next
	}

	if nil == currElem.Next {
		return nil, errors.New(nonValidPositionMsg)
	}

	newNode := &ListNode{Value: value, Next: currElem.Next}
	currElem.Next = newNode
	return list, nil
}

// Time complexity: O(n)
func (list *LinkedList) Length() int {
	if list.Empty() {
		return 0
	}

	sum := 1
	currElem := list.Head
	for currElem.Next != nil {
		sum++
		currElem = currElem.Next
	}
	return sum
}

// Time complexity: O(n)
func (list *LinkedList) Contains(value interface{}) bool {
	if list.Empty() {
		return false
	}

	currElem := list.Head
	for currElem != nil {
		// TODO: Check DeepEqual time complexity
		if reflect.DeepEqual(currElem.Value, value) {
			return true
		}
		currElem = currElem.Next
	}
	return false
}

// Time complexity: O(n)
func (list *LinkedList) OccurrencesOf(value interface{}) int {
	occurrences := 0
	if list.Empty() {
		return occurrences
	}

	currElem := list.Head
	for currElem != nil {
		// TODO: Check DeepEqual time complexity
		if reflect.DeepEqual(currElem.Value, value) {
			occurrences++
		}
		currElem = currElem.Next
	}
	return occurrences
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

// Time complexity: O(?)
func (list *LinkedList) String(strFn strings.Fn) string {
	var b bytes.Buffer // TODO: Pre-calculate the size of the resulting string
	var currNode *ListNode

	if nil != list.Head {
		b.WriteString(strFn(list.Head.Value))
		currNode = list.Head.Next
	}

	for nil != currNode {
		b.WriteString(" ")
		b.WriteString(strFn(currNode.Value))
		currNode = currNode.Next
	}
	return b.String()
}
