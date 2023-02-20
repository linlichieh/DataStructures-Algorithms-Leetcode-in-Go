package tree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		input []int
		root  *node // expected root and its children
	}{
		{
			input: []int{9},
			root:  &node{val: 9},
		},
		{
			input: []int{9, 4},
			root: &node{
				val:  9,
				left: &node{val: 4},
			},
		},
		{
			input: []int{9, 4, 6},
			root: &node{
				val: 9,
				left: &node{
					val:   4,
					right: &node{val: 6},
				},
			},
		},
		{
			input: []int{9, 4, 6, 20},
			root: &node{
				val: 9,
				left: &node{
					val:   4,
					right: &node{val: 6},
				},
				right: &node{val: 20},
			},
		},
		{
			input: []int{9, 4, 6, 20, 170},
			root: &node{
				val: 9,
				left: &node{
					val:   4,
					right: &node{val: 6},
				},
				right: &node{
					val:   20,
					right: &node{val: 170},
				},
			},
		},
		{
			input: []int{9, 4, 6, 20, 170, 15},
			root: &node{
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
		},
		{
			input: []int{9, 4, 6, 20, 170, 15, 1},
			root: &node{
				val: 9,
				left: &node{
					val:   4,
					left:  &node{val: 1},
					right: &node{val: 6},
				},
				right: &node{
					val:   20,
					left:  &node{val: 15},
					right: &node{val: 170},
				},
			},
		},
	}

	for i, tc := range testCases {
		tree := BST{}
		for _, val := range tc.input {
			tree.Insert(val)
		}

		if !reflect.DeepEqual(tc.root, tree.root) {
			t.Errorf("test case(%d): expect '%v', but got '%v", i, tc.root, tree.root)
		}
	}
}

func TestLookup(t *testing.T) {
	testCases := []struct {
		input []int
		val   int // the value to be looked up
		found bool
	}{
		{
			input: []int{},
			val:   9,
			found: false,
		},
		{
			input: []int{9},
			val:   9,
			found: true,
		},
		{
			input: []int{9, 4, 6},
			val:   6,
			found: true,
		},
		{
			input: []int{9, 4, 6},
			val:   5,
			found: false,
		},
		{
			input: []int{9, 4, 6, 20, 170},
			val:   170,
			found: true,
		},
		{
			input: []int{9, 4, 6, 20, 170},
			val:   4,
			found: true,
		},
		{
			input: []int{9, 4, 6, 20, 170},
			val:   171,
			found: false,
		},
		{
			input: []int{9, 4, 6, 20, 170, 15, 1},
			val:   15,
			found: true,
		},
		{
			input: []int{9, 4, 6, 20, 170, 15, 1},
			val:   1,
			found: true,
		},
		{
			input: []int{9, 4, 6, 20, 170, 15, 1},
			val:   4,
			found: true,
		},
		{
			input: []int{9, 4, 6, 20, 170, 15, 1},
			val:   16,
			found: false,
		},
	}

	for i, tc := range testCases {
		tree := BST{}
		for _, val := range tc.input {
			tree.Insert(val)
		}
		result := tree.Lookup(tc.val)
		if tc.found != result {
			t.Errorf("test case(%d): search: %d, expect '%v', but got '%v'", i, tc.val, tc.found, result)
		}
	}
}
