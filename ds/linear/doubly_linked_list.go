package linear

import (
	"bytes"
	"github.com/joanlopez/golang-hacker-rank/compare"
	"github.com/joanlopez/golang-hacker-rank/strings"
	"fmt"
)

const nonValidPositionMsg = "a wrong position was given: out of bounds"

type DoublyLinkedList struct {
	head   *DoublyLinkedNode
	tail   *DoublyLinkedNode
	length int
	cmpFn  compare.Fn
	strFn  strings.Fn
}

type DoublyLinkedNode struct {
	data interface{}
	next *DoublyLinkedNode
	prev *DoublyLinkedNode
}

func NewDoublyLinkedList(cmpFn compare.Fn, strFn strings.Fn) *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0, cmpFn, strFn}
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Empty() bool {
	return l.length == 0
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Length() int {
	return l.length
}

// Time complexity: O(n)
func (l *DoublyLinkedList) Contains(search interface{}) bool {
	currNode := l.head
	for nil != currNode {
		if 0 == l.cmpFn(search, currNode.Data()) {
			return true
		}
		currNode = currNode.Next()
	}
	return false
}

// Time complexity: O(n)
func (l *DoublyLinkedList) Occurrences(search interface{}) int {
	currNode := l.head
	occurrences := 0
	for nil != currNode {
		if 0 == l.cmpFn(search, currNode.Data()) {
			occurrences++
		}
		currNode = currNode.Next()
	}
	return occurrences
}

// Time complexity: O(n)
func (l *DoublyLinkedList) At(search interface{}) int {
	currNode := l.head
	pos := 0
	for nil != currNode {
		if 0 == l.cmpFn(search, currNode.Data()) {
			return pos
		}
		currNode = currNode.Next()
		pos++
	}
	return -1
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Prepend(data interface{}) {
	newNode := &DoublyLinkedNode{data, nil, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length = 1
		return
	}
	l.head.prepend(newNode)
	if l.length == 1 {
		l.tail = l.head
	}
	l.head = newNode
	l.length++
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Append(data interface{}) {
	newNode := &DoublyLinkedNode{data, nil, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length = 1
		return
	}
	l.tail.append(newNode)
	l.tail = newNode
	l.length++
}

// Time complexity: O(n)
func (l *DoublyLinkedList) InsertAt(data interface{}, pos int) error {
	if pos > l.length {
		return fmt.Errorf(nonValidPositionMsg)
	}
	if pos == 0 {
		l.Prepend(data)
		return nil
	}
	if pos == l.length {
		l.Append(data)
		return nil
	}
	currNode := l.head
	for pos > 0 {
		currNode = currNode.Next()
		pos--
	}
	currNode.prepend(&DoublyLinkedNode{data, nil, nil})
	l.length++
	return nil
}

// Time complexity: O(n)
func (l *DoublyLinkedList) GetAt(pos int) (interface{}, error) {
	if pos >= l.length {
		return nil, fmt.Errorf(nonValidPositionMsg)
	}
	if pos == 0 {
		return l.head.Data(), nil
	}
	if pos == l.length-1 {
		return l.tail.Data(), nil
	}
	currNode := l.head
	for pos > 0 {
		currNode = currNode.Next()
		pos--
	}
	return currNode.Data(), nil
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Shift() interface{} {
	if l.Empty() {
		return nil
	}
	elem := l.head
	l.head = elem.Next()
	if l.length == 1 {
		l.tail = nil
	}
	elem.unlink()
	l.length--
	return elem.Data()
}

// Time complexity: O(1)
func (l *DoublyLinkedList) Pop() interface{} {
	if l.Empty() {
		return nil
	}
	elem := l.tail
	l.tail = elem.Prev()
	if l.length == 1 {
		l.head = nil
	}
	elem.unlink()
	l.length--
	return elem.Data()
}

// Time complexity: O(n)
func (l *DoublyLinkedList) DeleteAt(pos int) (interface{}, error) {
	if pos >= l.length {
		return nil, fmt.Errorf(nonValidPositionMsg)
	}
	if pos == 0 {
		return l.Shift(), nil
	}
	if pos == l.length-1 {
		return l.Pop(), nil
	}
	currNode := l.head
	for pos > 0 {
		currNode = currNode.Next()
		pos--
	}
	currNode.unlink()
	l.length--
	return currNode.Data(), nil


}

// Time complexity: O(?)
func (l *DoublyLinkedList) String() string {
	if l.Empty() {
		return ""
	}
	var b bytes.Buffer
	b.WriteString(l.strFn(l.head.Data()))
	currNode := l.head.Next()
	for nil != currNode {
		b.WriteString(" ")
		b.WriteString(l.strFn(currNode.Data()))
		currNode = currNode.Next()
	}
	return b.String()
}

// Proxy method
func (l *DoublyLinkedList) Enqueue(data interface{}) {
	l.Append(data)
}

// Proxy method
func (l *DoublyLinkedList) Dequeue() interface{} {
	return l.Shift()
}

//
// Node functions
//

func (n *DoublyLinkedNode) Data() interface{} {
	if n == nil {
		return nil
	}
	return n.data
}

func (n *DoublyLinkedNode) Next() *DoublyLinkedNode {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *DoublyLinkedNode) Prev() *DoublyLinkedNode {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *DoublyLinkedNode) append(new *DoublyLinkedNode) {
	new.prev = n
	newNext := n.Next()
	new.next = newNext
	if newNext != nil {
		newNext.prev = new
	}
	n.next = new
}

func (n *DoublyLinkedNode) prepend(new *DoublyLinkedNode) {
	new.next = n
	newPrev := n.Prev()
	new.prev = newPrev
	if newPrev != nil {
		newPrev.next = new
	}
	n.prev = new
}

func (n *DoublyLinkedNode) unlink() {
	prev := n.Prev()
	next := n.Next()
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}