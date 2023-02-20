package tree

import (
	"reflect"
	"testing"
)

type traversalStruct struct {
	tree      *node
	traversal []int
}

func getBFSTeseCases() []traversalStruct {
	return []traversalStruct{
		{
			tree: &node{
				val: 9,
				left: &node{
					val:   4,
					right: &node{val: 6},
				},
				right: &node{
					val:   20,
					left:  &node{val: 15},
					right: &node{val: 170},
				},
			},
			traversal: []int{9, 4, 20, 6, 15, 170},
		},
		{
			tree: &node{
				val: 9,
				left: &node{
					val:   4,
					left:  &node{val: 3},
					right: &node{val: 6},
				},
				right: &node{
					val:   20,
					left:  &node{val: 15},
					right: &node{val: 170},
				},
			},

			traversal: []int{9, 4, 20, 3, 6, 15, 170},
		},
		{
			tree: &node{
				val: 9,
				left: &node{
					val:   4,
					right: &node{val: 6},
				},
				right: &node{
					val: 20,
					left: &node{
						val:   15,
						right: &node{val: 17},
					},
					right: &node{val: 170},
				},
			},
			traversal: []int{9, 4, 20, 6, 15, 170, 17},
		},
	}
}

func TestBFS(t *testing.T) {
	cases := getBFSTeseCases()

	for _, c := range cases {
		tree := BST{}
		tree.root = c.tree
		result := tree.BFS()
		if !reflect.DeepEqual(c.traversal, result) {
			t.Errorf("expect '%v', but got '%v'", c.traversal, result)
		}
	}
}

func TestBFSRecur(t *testing.T) {
	cases := getBFSTeseCases()

	for _, c := range cases {
		tree := BST{}
		tree.root = c.tree
		result := tree.BFSRecur([]*node{tree.root}, []int{})
		if !reflect.DeepEqual(c.traversal, result) {
			t.Errorf("expect '%v', but got '%v'", c.traversal, result)
		}
	}
}
