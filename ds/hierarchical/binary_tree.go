package hierarchical

import (
	"github.com/joanlopez/golang-hacker-rank/compare"
	"fmt"
	str "github.com/joanlopez/golang-hacker-rank/strings"
	"bytes"
	"github.com/joanlopez/golang-hacker-rank/ds/linear"
)

// Maxim number of nodes: 2^h - 1
// Maximum number of nodes at level ‘l’ = 2^(l-1)
// It's always balanced!
// Minimum possible height = ceil(Log2(n+1))
// # of leafs: # full nodes + 1
type BinaryTree struct {
	root  *BinaryTreeNode
	cmpFn compare.Fn
	strFn str.Fn
}

type BinaryTreeNode struct {
	data  interface{}
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func NewBinaryTree(cmpFn compare.Fn, strFn str.Fn) *BinaryTree {
	return &BinaryTree{nil, cmpFn, strFn}
}

// Time Complexity: O(n)
func (t *BinaryTree) Contains(pattern interface{}) bool {
	return t.root.contains(pattern, t.cmpFn)
}

// Time Complexity: O(log(n))
func (t *BinaryTree) Insert(data interface{}) {
	newNode := &BinaryTreeNode{data, nil, nil}
	if t.root == nil {
		t.root = newNode
		return
	}
	t.root.insert(newNode)
}

// Time Complexity: O(n)
func (t *BinaryTree) CountLeafs() int {
	if t.root == nil {
		return 0
	}
	return t.root.countFullChildren() + 1
}

// Time complexity: O(n)
func (t *BinaryTree) Delete(data interface{}) {

}

// Time complexity: O(n) without taking into account the strings management
func (t *BinaryTree) InOrder() string {
	return t.root.inOrder(t.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (t *BinaryTree) PreOrder() string {
	return t.root.preOrder(t.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (t *BinaryTree) PostOrder() string {
	return t.root.postOrder(t.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (t *BinaryTree) LevelOrder() string {
	if t.root == nil {
		return ""
	}
	queue := linear.NewQueue(t.strFn)
	var b bytes.Buffer
	b.WriteString(t.strFn(t.root.data))
	enqueueChildren(t.root, &queue)

	for true != queue.Empty() {
		next := queue.Dequeue().(*BinaryTreeNode)
		if next == nil {
			continue
		}
		b.WriteString(" ")
		b.WriteString(t.strFn(next.data))
		enqueueChildren(next, &queue)
	}
	return b.String()
}

//
// Node functions
//

func (n *BinaryTreeNode) contains(pattern interface{}, cmpFn compare.Fn) bool {
	if n == nil {
		return false
	}

	if 0 == cmpFn(n.data, pattern) {
		return true
	}

	foundOnLeft := n.left.contains(pattern, cmpFn)
	foundOnRight := n.right.contains(pattern, cmpFn)

	return foundOnLeft || foundOnRight
}

func (n *BinaryTreeNode) insert(newNode *BinaryTreeNode) {
	// TODO: Improve code
	if n.left == nil {
		n.left = newNode
		return
	}
	if n.right == nil {
		n.right = newNode
		return
	}
	if !n.left.isFull() {
		n.left.insert(newNode)
		return
	}
	if !n.right.isFull() {
		n.right.insert(newNode)
		return
	}
	n.left.insert(newNode)
}

func (n *BinaryTreeNode) isFull() bool {
	if n == nil {
		return false
	}
	return n.left != nil && n.right != nil
}

func (n *BinaryTreeNode) countFullChildren() int {
	if n == nil {
		return 0
	}

	childrenCount := n.left.countFullChildren() + n.right.countFullChildren()
	if n.isFull() {
		return childrenCount + 1
	}
	return childrenCount
}

func (n *BinaryTreeNode) swap(n2 *BinaryTreeNode) error {
	if n == nil || n2 == nil {
		return fmt.Errorf("swap error: one of the given nodes is nil")
	}

	tmp := n2.copy()
	assign(n, n2)
	assign(tmp, n)
	return nil
}

func (n *BinaryTreeNode) copy() *BinaryTreeNode {
	return &BinaryTreeNode{n.data, n.left, n.right}
}

func (n *BinaryTreeNode) inOrder(strFn str.Fn) string {
	if n == nil {
		return ""
	}
	return str.JoinByIgnoring([]string{
		n.left.inOrder(strFn),
		strFn(n.data),
		n.right.inOrder(strFn),
	}, " ", "")
}

func (n *BinaryTreeNode) preOrder(strFn str.Fn) string {
	if n == nil {
		return ""
	}
	return str.JoinByIgnoring([]string{
		strFn(n.data),
		n.left.preOrder(strFn),
		n.right.preOrder(strFn),
	}, " ", "")
}

func (n *BinaryTreeNode) postOrder(strFn str.Fn) string {
	if n == nil {
		return ""
	}
	return str.JoinByIgnoring([]string{
		n.left.postOrder(strFn),
		n.right.postOrder(strFn),
		strFn(n.data),
	}, " ", "")
}

func assign(from, to *BinaryTreeNode) {
	to.data = from.data
	to.left = from.left
	to.right = from.right
}

func enqueueChildren(n *BinaryTreeNode, q *linear.Queue) {
	(*q).Enqueue(n.left)
	(*q).Enqueue(n.right)
}
