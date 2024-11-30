package binarysearchtree

import (
	"sort"
)

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i, left: nil, right: nil}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i > bst.data {
		if bst.right == nil {
			bst.right = NewBst(i)
			return
		}
		bst.right.Insert(i)
		return
	}

	if bst.left == nil {
		bst.left = NewBst(i)
		return
	}

	bst.left.Insert(i)
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	type Node struct {
		who *BinarySearchTree
		lr bool
		rr bool
		added bool
	}
	nodes:=make([]*Node, 0)
	appendOnList:=func(bst *BinarySearchTree) {
		nodes = append(nodes, &Node{who: bst, lr: false, rr: false, added: false})
	}

	appendOnList(bst)
	result:=make([]int, 0)

	for len(nodes) > 0 {
		node:=nodes[len(nodes) - 1]
		if !node.added {
			result = append(result, node.who.data)
			node.added=true
		}

		if (!node.lr && node.who.left != nil) {
			node.lr=true
			appendOnList(node.who.left)
			continue
		}

		if !node.rr && node.who.right != nil {
			node.rr=true
			appendOnList(node.who.right)
			continue
		}
		
		nodes = append([]*Node{}, nodes[0:len(nodes)-1]...)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
