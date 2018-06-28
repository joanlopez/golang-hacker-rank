package linear

import (
	"testing"
)

// TODO: Refactor to table test approach with better messages & test cases

func TestLinkedListLenWithoutElements(t *testing.T) {
	var list *ListNode = nil
	got := list.Len()
	if 0 != got {
		t.Errorf("Error calculating the length of an empty list")
	}
}

func TestLinkedListLenWithOneElement(t *testing.T) {
	list := CreateListNode(1)
	got := list.Len()
	if 1 != got {
		t.Errorf("Error calculating the length of one element list")
	}
}

func TestLinkedListLenWithTwoElements(t *testing.T) {
	list := CreateListNode(1)
	list.Add(CreateListNode(2))
	got := list.Len()
	if 2 != got {
		t.Errorf("Error calculating the length of two elements list")
	}
}

func TestLinkedListContainsWithoutElements(t *testing.T) {
	var list *ListNode = nil
	got := list.Contains(1)
	if false != got {
		t.Errorf("Error checking if exists in an empty list")
	}
}

func TestLinkedListContainsWithOneElement(t *testing.T) {
	list := CreateListNode(1)
	got := list.Contains(1)
	if true != got {
		t.Errorf("Error checking if exists in a one element list")
	}
}

func TestLinkedListContainsWithTwoElements(t *testing.T) {
	list := CreateListNode(1)
	list.Add(CreateListNode(2))
	got := list.Contains(3)
	if false != got {
		t.Errorf("Error checking if exists in a two elements list")
	}
}

func TestLinkedListGetAtWithoutElements(t *testing.T) {
	var list *ListNode = nil
	got := list.GetAt(1)
	if nil != got {
		t.Errorf("Error getting at in an empty list")
	}
}

func TestLinkedListGetWithOneElement(t *testing.T) {
	list := CreateListNode(1)
	got := list.GetAt(0)
	if 1 != got {
		t.Errorf("Error getting at in a one element list")
	}
}

func TestLinkedListGetAtWithTwoElements(t *testing.T) {
	list := CreateListNode(1)
	list.Add(CreateListNode(2))
	got := list.GetAt(2)
	if nil != got {
		t.Errorf("Error getting at in a two elements list")
	}
}