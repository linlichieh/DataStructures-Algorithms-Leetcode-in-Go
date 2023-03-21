package leetcode133

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	cases := []struct {
		elements []int
		conns    [][]int
	}{
		{
			elements: []int{1, 2, 3, 4},
			conns:    [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			elements: []int{1},
			conns:    [][]int{},
		},
		{
			elements: []int{1, 2, 3, 4, 5, 6, 7},
			conns:    [][]int{{2, 7}, {1, 6}, {6, 4}, {3, 5}, {4, 6}, {2, 7, 3, 5}, {1, 6}},
		},
	}

	for i, c := range cases {
		// Build the graph
		nodes := []*Node{}    // input
		expected := []*Node{} // expected result (different addr from input because it's a cloned one)
		for _, v := range c.elements {
			nodes = append(nodes, &Node{Val: v})
			expected = append(expected, &Node{Val: v})
		}
		for i, vals := range c.conns {
			for _, val := range vals {
				// each node's value is the same as the node's index (1-indexed), so do `val-1`
				nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[val-1])
				expected[i].Neighbors = append(expected[i].Neighbors, expected[val-1])
			}
		}

		result := cloneGraph(nodes[0])
		if !reflect.DeepEqual(expected[0], result) || expected[0] == result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, expected, result)
		}
	}
}
