package hierarchical

import (
	"github.com/joanlopez/golang-hacker-rank/compare"
	"fmt"
	str "github.com/joanlopez/golang-hacker-rank/strings"
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
func (b *BinaryTree) Contains(pattern interface{}) bool {
	return b.root.contains(pattern, b.cmpFn)
}

// Time Complexity: O(log(n))
func (b *BinaryTree) Insert(data interface{}) {
	newNode := &BinaryTreeNode{data, nil, nil}
	if b.root == nil {
		b.root = newNode
		return
	}
	b.root.insert(newNode)
}

// Time Complexity: O(n)
func (b *BinaryTree) CountLeafs() int {
	if b.root == nil {
		return 0
	}
	return b.root.countFullChildren() + 1
}

// Time complexity: O(n)
func (b *BinaryTree) Delete(data interface{}) {

}

// Time complexity: O(n) without taking into account the strings management
func (b *BinaryTree) InOrder() string {
	return b.root.inOrder(b.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (b *BinaryTree) PreOrder() string {
	return b.root.preOrder(b.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (b *BinaryTree) PostOrder() string {
	return b.root.postOrder(b.strFn)
}

// Time complexity: O(n) without taking into account the strings management
func (b *BinaryTree) LevelOrder() string {
	return b.root.levelOrder(b.strFn)
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

func (n *BinaryTreeNode) levelOrder(strFn str.Fn) string {
	if n == nil {
		return ""
	}
	return str.JoinByIgnoring([]string{
		n.left.inOrder(strFn),
		strFn(n.data),
		n.right.inOrder(strFn),
	}, " ", "")
}

func assign(from, to *BinaryTreeNode) {
	to.data = from.data
	to.left = from.left
	to.right = from.right
}
