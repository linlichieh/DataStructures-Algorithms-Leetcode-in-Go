package leetcode98

import "testing"

type testCase struct {
	tree    *TreeNode
	isValid bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			tree:    &TreeNode{},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 7},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 3},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 7},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 7},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 7},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 8},
				},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   7,
					Right: &TreeNode{Val: 6},
				},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 8},
				},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:  7,
						Left: &TreeNode{Val: 6},
					},
					Right: &TreeNode{Val: 11},
				},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val:   5,
							Right: &TreeNode{Val: 6},
						},
					},
					Right: &TreeNode{Val: 11},
				},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 18,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val:   5,
							Right: &TreeNode{Val: 8},
						},
					},
					Right: &TreeNode{Val: 11},
				},
			},
			isValid: false,
		},
		// from leetcode
		{
			tree: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
			},
			isValid: false,
		},
		// from leetcode
		{
			tree: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			isValid: false,
		},
		{
			tree: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 12},
					},
				},
			},
			isValid: true,
		},
		{
			tree: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
				Right: &TreeNode{Val: 10},
			},
			isValid: false,
		},
		// from leetcode
		{
			tree: &TreeNode{
				Val: 120,
				Left: &TreeNode{
					Val: 70,
					Left: &TreeNode{
						Val:   50,
						Left:  &TreeNode{Val: 20},
						Right: &TreeNode{Val: 55},
					},
					Right: &TreeNode{
						Val:   100,
						Left:  &TreeNode{Val: 75},
						Right: &TreeNode{Val: 110},
					},
				},
				Right: &TreeNode{
					Val: 140,
					Left: &TreeNode{
						Val:   130,
						Left:  &TreeNode{Val: 119},
						Right: &TreeNode{Val: 135},
					},
					Right: &TreeNode{
						Val:   160,
						Left:  &TreeNode{Val: 150},
						Right: &TreeNode{Val: 200},
					},
				},
			},
			isValid: false,
		},
		{
			tree:    &TreeNode{Val: 2147483647},
			isValid: true,
		},
	}
}

func TestIsValid1(t *testing.T) {
	cases := getTestCases()
	for i, c := range cases {
		isValid := isValidBST1(c.tree)
		if isValid != c.isValid {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.isValid, isValid)
		}
	}
}
