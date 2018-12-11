package linear

import (
	"github.com/joanlopez/golang-hacker-rank/compare"
	"github.com/joanlopez/golang-hacker-rank/strings"
	"testing"
)

// TODO: Rewrite tests to improve the code

func TestEmptyMethodWhenEmpty(t *testing.T) {
	list := NewDoublyLinkedList(nil, nil)
	if false == list.Empty() {
		t.Errorf("Error on Empty() method")
	}
}

func TestEmptyMethodWhenHasItems(t *testing.T) {
	list := buildLinkedList()
	if true == list.Empty() {
		t.Errorf("Error on Empty() method")
	}
}

func TestLengthMethodWhenEmpty(t *testing.T) {
	list := NewDoublyLinkedList(nil, nil)
	if 0 != list.Length() {
		t.Errorf("Error on Length() method")
	}
}

func TestLengthMethodWhenHasItems(t *testing.T) {
	list := buildLinkedList()
	if 5 != list.Length() {
		t.Errorf("Error on Length() method")
	}
}

func TestContainsMethodWhenTrue(t *testing.T) {
	list := buildLinkedList()
	if true != list.Contains(2) {
		t.Errorf("Error on Contains() method")
	}
}

func TestContainsMethodWhenFalse(t *testing.T) {
	list := buildLinkedList()
	if false != list.Contains(1) {
		t.Errorf("Error on Contains() method")
	}
}

func TestOccurrencesWhenContains(t *testing.T) {
	list := buildLinkedList()
	if 1 != list.Occurrences(2) {
		t.Errorf("Error on Occurrences() method")
	}
}

func TestOccurrencesWhenNoContains(t *testing.T) {
	list := buildLinkedList()
	if 0 != list.Occurrences(1) {
		t.Errorf("Error on Occurrences() method")
	}
}

func TestAtWhenContains(t *testing.T) {
	list := buildLinkedList()
	if 0 != list.At(2) {
		t.Errorf("Error on At() method")
	}
}

func TestAtWhenNoContains(t *testing.T) {
	list := buildLinkedList()
	if -1 != list.At(1) {
		t.Errorf("Error on At() method")
	}
}

// TODO: Ugly test ;) ;)
func TestInsertGetDeleteSequence(t *testing.T) {
	errMsg := "Error on InsertGetDelete sequence"
	list := buildLinkedList()
	if false != list.Contains(5) {
		t.Errorf(errMsg)
	}
	err := list.InsertAt(5, 2)
	if err != nil {
		t.Errorf(errMsg)
	}
	if v, _ := list.GetAt(2); 0 != compare.IntFn(5, v) {
		t.Errorf(errMsg)
	}
	if true != list.Contains(5) {
		t.Errorf(errMsg)
	}
	if v, _ := list.DeleteAt(2); 0 != compare.IntFn(5, v) {
		t.Errorf(errMsg)
	}
	if false != list.Contains(5) {
		t.Errorf(errMsg)
	}
}

func TestShiftWhenContains(t *testing.T) {
	errMsg := "Error on Shift() method"
	list := buildLinkedList()
	if 2 != list.Shift() {
		t.Errorf(errMsg)
	}
	if 4 != list.Length() {
		t.Errorf(errMsg)
	}
}

func TestShiftWhenNoContains(t *testing.T) {
	list := NewDoublyLinkedList(nil, nil)
	if nil != list.Shift() {
		t.Errorf("Error on Shift() method")
	}
}

func TestPopWhenContains(t *testing.T) {
	errMsg := "Error on Pop() method"
	list := buildLinkedList()
	if 10 != list.Pop() {
		t.Errorf(errMsg)
	}
	if 4 != list.Length() {
		t.Errorf(errMsg)
	}
}

func TestPopWhenNoContains(t *testing.T) {
	list := NewDoublyLinkedList(nil, nil)
	if nil != list.Pop() {
		t.Errorf("Error on Pop() method")
	}
}

func TestStringWhenNotEmpty(t *testing.T) {
	list := buildLinkedList()
	if "2 4 6 8 10" != list.String() {
		t.Errorf("Error on String() method")
	}
}

func TestStringWhenEmpty(t *testing.T) {
	list := NewDoublyLinkedList(nil, nil)
	if "" != list.String() {
		t.Errorf("Error on String() method")
	}
}

// Builds the following linked list:
//
// 2 - 4 - 6 - 8 - 10
//
func buildLinkedList() *DoublyLinkedList{
	list := NewDoublyLinkedList(compare.IntFn, strings.IntFn)
	list.Append(4)
	list.Append(6)
	list.Append(8)
	list.Append(10)
	list.Prepend(2)
	return list
}