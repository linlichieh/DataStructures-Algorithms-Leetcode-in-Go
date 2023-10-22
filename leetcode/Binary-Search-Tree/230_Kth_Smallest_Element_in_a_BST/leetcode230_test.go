package leetcode230

import "testing"

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		root     *TreeNode
		k        int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			k:        1,
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:        3,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 5},
			},
			k:        3,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			k:        3,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			k:        2,
			expected: 3,
		},
	}

	for i, c := range cases {
		result := kthSmallest(c.root, c.k)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
