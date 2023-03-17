package leetcode297

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		input  *TreeNode
		output string
	}{
		{
			input: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			output: "[1,2,null,null,3,4,null,null,5,null,null]",
		},
		{
			input:  nil,
			output: "[]",
		},
	}

	for i, c := range cases {
		ser := Constructor()
		result := ser.serialize(c.input)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}

func TestDeserialize(t *testing.T) {
	cases := []struct {
		input  string
		output *TreeNode
	}{
		{
			input: "[1,2,null,null,3,4,null,null,5,null,null]",
			output: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			input:  "[]",
			output: nil,
		},
	}

	for i, c := range cases {
		deser := Constructor()
		result := deser.deserialize(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
