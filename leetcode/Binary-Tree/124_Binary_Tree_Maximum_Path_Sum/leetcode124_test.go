package leetcode124

import (
	"reflect"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		output int
	}{
		{
			root:   &TreeNode{},
			output: 0,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			output: 6,
		},
		{
			root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			output: 42,
		},
		{
			root:   &TreeNode{Val: 1},
			output: 1,
		},
		{
			root: &TreeNode{
				Val:  -2,
				Left: &TreeNode{Val: -1},
			},
			output: -1,
		},
	}

	for i, c := range cases {
		result := maxPathSum(c.root)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
