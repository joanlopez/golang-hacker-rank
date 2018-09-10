package hierarchical

import (
	"testing"
	"github.com/joanlopez/golang-hacker-rank/compare"
	"github.com/joanlopez/golang-hacker-rank/strings"
)

type TraverseFn func(b *BinaryTree) string

type TraverseTC struct {
	title    string
	expected string
	fn       TraverseFn
}

var tests = []TraverseTC{
	{"inOrder traversing", "4 2 5 1 3", func(b *BinaryTree) string { return b.InOrder() }},
	{"preOrder traversing", "1 2 4 5 3", func(b *BinaryTree) string { return b.PreOrder() }},
	{"postOrder traversing", "4 5 2 3 1", func(b *BinaryTree) string { return b.PostOrder() }},
	{"levelOrder traversing", "1 2 3 4 5", func(b *BinaryTree) string { return b.LevelOrder() }},
}

func TestBinaryTreeTraverses(t *testing.T) {
	tree := buildTestingTree()
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			got := tc.fn(tree)
			if tc.expected != got {
				t.Errorf("Error while %v; expected: %v - got: %v", tc.title, tc.expected, got)
			}
		})
	}
}

// Builds the following binary tree:
//
//				1
//		2	 			3
//	4		5		-		-
//
func buildTestingTree() *BinaryTree {
	tree := NewBinaryTree(compare.IntFn, strings.IntFn)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	return tree
}
