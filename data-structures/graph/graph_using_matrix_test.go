package graph

import (
	"reflect"
	"testing"
)

func TestMatrixAddEdge(t *testing.T) {
	testCases := []struct {
		num       int
		edges     [][]int
		adjMatrix [][]int
	}{
		{
			num:       0,
			edges:     [][]int{},
			adjMatrix: [][]int{},
		},
		{
			num:   1,
			edges: [][]int{},
			adjMatrix: [][]int{
				[]int{0},
			},
		},
		{
			num:   2,
			edges: [][]int{},
			adjMatrix: [][]int{
				[]int{0, 0},
				[]int{0, 0},
			},
		},
		{
			num: 2,
			edges: [][]int{
				[]int{0, 1},
			},
			adjMatrix: [][]int{
				[]int{0, 1},
				[]int{1, 0},
			},
		},
		{
			num: 4,
			edges: [][]int{
				[]int{1, 2},
				[]int{0, 3},
			},
			adjMatrix: [][]int{
				[]int{0, 0, 0, 1},
				[]int{0, 0, 1, 0},
				[]int{0, 1, 0, 0},
				[]int{1, 0, 0, 0},
			},
		},
		{
			num: 8,
			edges: [][]int{
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 5},
				[]int{4, 5},
				[]int{6, 7},
			},
			adjMatrix: [][]int{
				[]int{0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 1, 1, 0, 0, 0},
				[]int{0, 0, 1, 0, 0, 1, 0, 0},
				[]int{0, 0, 1, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 1, 1, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1},
				[]int{0, 0, 0, 0, 0, 0, 1, 0},
			},
		},
	}

	for _, tc := range testCases {
		g := NewGraphMatrix(tc.num)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}
		if !reflect.DeepEqual(tc.adjMatrix, g.adjacencyMatrix) {
			t.Errorf("expect '%v', but got '%v'", tc.adjMatrix, g.adjacencyMatrix)
		}
	}
}

func TestMatrixGetNeighbors(t *testing.T) {
	type expected struct {
		v         int
		neighbors []int
	}

	testCases := []struct {
		num      int
		edges    [][]int
		expected []expected
	}{
		{
			num:   6,
			edges: [][]int{},
			expected: []expected{
				{2, []int{}},
			},
		},
		{
			num:   8,
			edges: [][]int{{5, 3}, {5, 4}, {3, 2}, {2, 4}, {6, 7}},
			expected: []expected{
				{2, []int{3, 4}},
				{3, []int{2, 5}},
				{5, []int{3, 4}},
				{7, []int{6}},
			},
		},
	}

	for i, tc := range testCases {
		g := NewGraphMatrix(tc.num)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		for j, l := range tc.expected {
			result := g.GetNeighbors(l.v)
			if !reflect.DeepEqual(l.neighbors, result) {
				t.Errorf("(%d-%d) expect '%v', but got '%v'", i, j, l.neighbors, result)
			}
		}
	}
}

func TestMatrixIsAdjacent(t *testing.T) {
	type expected struct {
		uv    []int // represent vertex u and vertex v
		isAdj bool
	}

	testCases := []struct {
		num      int
		edges    [][]int
		expected []expected
	}{
		{
			num:   6,
			edges: [][]int{},
			expected: []expected{
				{[]int{2, 3}, false},
			},
		},
		{
			num:   8,
			edges: [][]int{{5, 3}, {5, 4}, {3, 2}, {2, 4}, {6, 7}},
			expected: []expected{
				{[]int{2, 3}, true},
				{[]int{3, 5}, true},
				{[]int{5, 2}, false},
				{[]int{2, 6}, false},
				{[]int{4, 2}, true},
				{[]int{6, 7}, true},
				{[]int{5, 4}, true},
				{[]int{3, 4}, false},
				{[]int{7, 4}, false},
			},
		},
	}

	for _, tc := range testCases {
		g := NewGraphMatrix(tc.num)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		for _, l := range tc.expected {
			result := g.IsAdjacent(l.uv[0], l.uv[1])
			if result != l.isAdj {
				t.Errorf("expect '%v', but got '%v'", l.isAdj, result)
			}
		}
	}
}
