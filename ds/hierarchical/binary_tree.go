package hierarchical

import "golang-hacker-rank/compare"

// Maxim number of nodes: 2^h - 1
// Maximum number of nodes at level ‘l’ = 2^(l-1)
// It's always balanced!
// Minimum possible height = ceil(Log2(n+1))
// # of leafs: # full nodes + 1
type BinaryTree struct {
	root *BinaryTreeNode
	cmpFn compare.Fn
}

type BinaryTreeNode struct {
	data interface{}
	left *BinaryTreeNode
	right *BinaryTreeNode
}

func NewBinaryTree(cmpFn compare.Fn) BinaryTree {
	return BinaryTree{nil, cmpFn}
}

// Time Complexity: O(n)
func (b BinaryTree) Contains(pattern interface{}) bool {
	return b.root.contains(pattern, b.cmpFn)
}

// Time Complexity: O(log(n))
func (b BinaryTree) Insert(data interface{}) {
	newNode := &BinaryTreeNode{data, nil, nil}
	if b.root != nil {
		b.root.insert(newNode)
	}
	b.root = newNode
}

// Time Complexity: O(n)
func (b BinaryTree) Leafs() int {
	if b.root == nil {
		return 0
	}
	return b.root.countFullChildren() + 1
}

func (b *BinaryTreeNode) contains(pattern interface{}, cmpFn compare.Fn) bool {
	if b == nil {
		return false
	}

	if 0 == cmpFn(b.data, pattern) {
		return true
	}

	foundOnLeft := b.left.contains(pattern, cmpFn)
	foundOnRight := b.right.contains(pattern, cmpFn)

	return foundOnLeft || foundOnRight
}

func (b *BinaryTreeNode) insert(newNode *BinaryTreeNode) {
	if b.left == nil {
		b.left = newNode
		return
	}
	if b.right == nil {
		b.right = newNode
		return
	}
	if !b.left.isFull() {
		b.left.insert(newNode)
	}
	if !b.right.isFull() {
		b.right.insert(newNode)
	}
	b.left.insert(newNode)
}

func (b *BinaryTreeNode) isFull() bool {
	if b == nil {
		return false
	}
	return b.left != nil && b.right != nil
}

func (b *BinaryTreeNode) countFullChildren() int {
	if b == nil {
		return 0
	}

	childrenCount := b.left.countFullChildren() + b.right.countFullChildren()
	if b.isFull() {
		return childrenCount + 1
	}
	return childrenCount
}