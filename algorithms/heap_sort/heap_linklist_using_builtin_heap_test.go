package heapSort

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	cases2 := []struct {
		input  *Node
		output *Node
	}{
		{
			input:  &Node{Val: 2, Next: &Node{Val: 1, Next: &Node{Val: 5, Next: &Node{Val: 3}}}},
			output: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 5}}}},
		},
		{
			input:  &Node{Val: 53, Next: &Node{Val: 11, Next: &Node{Val: 38, Next: &Node{Val: 13, Next: &Node{Val: 71, Next: &Node{Val: 2, Next: &Node{Val: 31, Next: &Node{Val: 23}}}}}}}},
			output: &Node{Val: 2, Next: &Node{Val: 11, Next: &Node{Val: 13, Next: &Node{Val: 23, Next: &Node{Val: 31, Next: &Node{Val: 38, Next: &Node{Val: 53, Next: &Node{Val: 71}}}}}}}},
		},
	}

	for i, c := range cases2 {
		result := sortList(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
