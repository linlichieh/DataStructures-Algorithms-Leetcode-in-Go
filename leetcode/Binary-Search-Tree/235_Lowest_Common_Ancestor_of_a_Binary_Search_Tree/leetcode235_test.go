package leetcode235

import (
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		output *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 2},
			q: &TreeNode{Val: 8},
			output: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 2},
			q: &TreeNode{Val: 4},
			output: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			p: &TreeNode{Val: 2},
			q: &TreeNode{Val: 1},
			output: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 5},
			q: &TreeNode{Val: 7},
			output: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 7},
			q: &TreeNode{Val: 9},
			output: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 7},
			q: &TreeNode{Val: 6},
			output: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p: &TreeNode{Val: 0},
			q: &TreeNode{Val: 5},
			output: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
		},
	}

	for i, c := range cases {
		result := lowestCommonAncestor(c.root, c.p, c.q)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
