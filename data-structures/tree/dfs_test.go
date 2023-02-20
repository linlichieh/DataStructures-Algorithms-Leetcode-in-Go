package tree

import (
	"reflect"
	"testing"
)

func TestDFSInOrder(t *testing.T) {
	cases := []struct {
		tree      *node
		traversal []int
	}{
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
			traversal: []int{4, 6, 9, 15, 20, 170},
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

			traversal: []int{3, 4, 6, 9, 15, 20, 170},
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
			traversal: []int{4, 6, 9, 15, 17, 20, 170},
		},
	}

	for _, c := range cases {
		tree := BST{}
		tree.root = c.tree
		result := tree.DFSInOrder()
		if !reflect.DeepEqual(c.traversal, result) {
			t.Errorf("expect '%v', but got '%v'", c.traversal, result)
		}
	}
}

func TestDFSPreOrder(t *testing.T) {
	cases := []struct {
		tree      *node
		traversal []int
	}{
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
			traversal: []int{9, 4, 6, 20, 15, 170},
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

			traversal: []int{9, 4, 3, 6, 20, 15, 170},
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
			traversal: []int{9, 4, 6, 20, 15, 17, 170},
		},
	}

	for _, c := range cases {
		tree := BST{}
		tree.root = c.tree
		result := tree.DFSPreOrder()
		if !reflect.DeepEqual(c.traversal, result) {
			t.Errorf("expect '%v', but got '%v'", c.traversal, result)
		}
	}
}

func TestDFSPostOrder(t *testing.T) {
	cases := []struct {
		tree      *node
		traversal []int
	}{
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
			traversal: []int{6, 4, 15, 170, 20, 9},
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

			traversal: []int{3, 6, 4, 15, 170, 20, 9},
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
			traversal: []int{6, 4, 17, 15, 170, 20, 9},
		},
	}

	for _, c := range cases {
		tree := BST{}
		tree.root = c.tree
		result := tree.DFSPostOrder()
		if !reflect.DeepEqual(c.traversal, result) {
			t.Errorf("expect '%v', but got '%v'", c.traversal, result)
		}
	}
}
