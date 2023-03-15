package leetcode105

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	cases := []struct {
		preorder []int
		inorder  []int
		tree     *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			tree: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			tree:     &TreeNode{Val: -1},
		},
		{
			preorder: []int{7, 5, 8, 6, 4, 13, 21, 9},
			inorder:  []int{8, 5, 4, 6, 7, 13, 9, 21},
			tree: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 8},
					Right: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val:  21,
						Left: &TreeNode{Val: 9},
					},
				},
			},
		},
	}

	for i, c := range cases {
		result := buildTree(c.preorder, c.inorder)
		if !reflect.DeepEqual(c.tree, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.tree, result)
		}
	}
}
