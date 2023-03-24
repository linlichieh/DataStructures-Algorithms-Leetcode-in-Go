package leetcode226

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	cases := []struct {
		input  *TreeNode
		output *TreeNode
	}{
		{
			input:  &TreeNode{},
			output: &TreeNode{},
		},
		{
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			output: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}

	for i, c := range cases {
		result := invertTree(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
