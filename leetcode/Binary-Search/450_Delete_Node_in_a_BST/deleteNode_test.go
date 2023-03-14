package leetcode450

import (
	"reflect"
	"testing"
)

type testCase struct {
	comment      string
	originalTree *TreeNode // tree
	val          int       // the Value to be removed
	updatedTree  *TreeNode // trimmed tree
}

// For both recursive and non-recursive implementation
func getTestCases() []testCase {
	return []testCase{
		{
			comment:      "Delete an node from empty tree",
			originalTree: &TreeNode{},
			val:          9,
			updatedTree:  &TreeNode{},
		},
		{
			comment: "Delete an inexistent node",
			originalTree: &TreeNode{
				Val:   100,
				Left:  &TreeNode{Val: 75},
				Right: &TreeNode{Val: 125},
			},
			val: 88,
			updatedTree: &TreeNode{
				Val:   100,
				Left:  &TreeNode{Val: 75},
				Right: &TreeNode{Val: 125},
			},
		},
		{
			comment: "case 1: Delete root node with no child",
			originalTree: &TreeNode{
				Val: 100,
			},
			val:         100,
			updatedTree: nil,
		},
		{
			comment: "case 1: Delete a leaf (no child)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 70,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:  65,
						Left: &TreeNode{Val: 60},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 1: Delete a leaf (no child)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 115,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 2: Delete root node with left child only",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
			},
			val: 100,
			updatedTree: &TreeNode{
				Val: 75,
				Left: &TreeNode{
					Val:   65,
					Left:  &TreeNode{Val: 60},
					Right: &TreeNode{Val: 70},
				},
				Right: &TreeNode{
					Val:   85,
					Left:  &TreeNode{Val: 80},
					Right: &TreeNode{Val: 95},
				},
			},
		},
		{
			comment: "case 2: Delete root node with right child only",
			originalTree: &TreeNode{
				Val: 100,
				Right: &TreeNode{
					Val:  125,
					Left: &TreeNode{Val: 115},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 100,
			updatedTree: &TreeNode{
				Val:  125,
				Left: &TreeNode{Val: 115},
				Right: &TreeNode{
					Val:   150,
					Left:  &TreeNode{Val: 135},
					Right: &TreeNode{Val: 175},
				},
			},
		},
		{
			comment: "case 2: Delete a node with left child only",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
				},
			},
			val: 75,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val:   65,
					Left:  &TreeNode{Val: 60},
					Right: &TreeNode{Val: 70},
				},
			},
		},
		{
			comment: "case 2: Delete a node with left child only",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:  65,
						Left: &TreeNode{Val: 60},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
			},
			val: 65,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val: 60,
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
			},
		},
		{
			comment: "case 2: Delete a node with right child only",
			originalTree: &TreeNode{
				Val: 100,
				Right: &TreeNode{
					Val: 125,
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 125,
			updatedTree: &TreeNode{
				Val: 100,
				Right: &TreeNode{
					Val:   150,
					Left:  &TreeNode{Val: 135},
					Right: &TreeNode{Val: 175},
				},
			},
		},
		{
			comment: "case 2: Delete a node with right child only",
			originalTree: &TreeNode{
				Val: 100,
				Right: &TreeNode{
					Val:  125,
					Left: &TreeNode{Val: 115},
					Right: &TreeNode{
						Val:   150,
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 150,
			updatedTree: &TreeNode{
				Val: 100,
				Right: &TreeNode{
					Val:   125,
					Left:  &TreeNode{Val: 115},
					Right: &TreeNode{Val: 175},
				},
			},
		},
		{
			comment: "case 2: Delete a node with right child only (from leetcode)",
			originalTree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			val: 1,
			updatedTree: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
		},
		{
			comment: "case 2: Delete a node with left child only (from leetcode)",
			originalTree: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
				},
			},
			val: 5,
			updatedTree: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4},
			},
		},
		{
			comment: "case 3: Delete a node with 2 children and leftmost node has right node and its children (from leetcode)",
			originalTree: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:  33,
					Left: &TreeNode{Val: 25},
					Right: &TreeNode{
						Val: 40,
						Left: &TreeNode{
							Val: 34,
							Right: &TreeNode{
								Val:   36,
								Left:  &TreeNode{Val: 35},
								Right: &TreeNode{Val: 39},
							},
						},
						Right: &TreeNode{Val: 45},
					},
				},
			},
			val: 33,
			updatedTree: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:  34,
					Left: &TreeNode{Val: 25},
					Right: &TreeNode{
						Val: 40,
						Left: &TreeNode{
							Val:   36,
							Left:  &TreeNode{Val: 35},
							Right: &TreeNode{Val: 39},
						},
						Right: &TreeNode{Val: 45},
					},
				},
			},
		},
		{
			comment: "case 3: Delete root node with 2 children",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 100,
			updatedTree: &TreeNode{
				Val: 115,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 3: Delete left node with 2 children (1)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 75,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 80,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 3: Delete left node with 2 children (2)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 65,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:  70,
						Left: &TreeNode{Val: 60},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 3: Delete right node with 2 children (1)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 125,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 135,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Right: &TreeNode{Val: 175},
					},
				},
			},
		},
		{
			comment: "case 3: Delete right node with 2 children (2)",
			originalTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:   150,
						Left:  &TreeNode{Val: 135},
						Right: &TreeNode{Val: 175},
					},
				},
			},
			val: 150,
			updatedTree: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
					Left: &TreeNode{
						Val:   65,
						Left:  &TreeNode{Val: 60},
						Right: &TreeNode{Val: 70},
					},
					Right: &TreeNode{
						Val:   85,
						Left:  &TreeNode{Val: 80},
						Right: &TreeNode{Val: 95},
					},
				},
				Right: &TreeNode{
					Val: 125,
					Left: &TreeNode{
						Val: 115,
					},
					Right: &TreeNode{
						Val:  175,
						Left: &TreeNode{Val: 135},
					},
				},
			},
		},
	}
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range getTestCases() {
		returnedTree := deleteNode(tc.originalTree, tc.val)
		if !reflect.DeepEqual(tc.updatedTree, returnedTree) {
			t.Errorf("%s - expect '%v', but got '%v'", tc.comment, tc.updatedTree, returnedTree)
		}
	}
}
