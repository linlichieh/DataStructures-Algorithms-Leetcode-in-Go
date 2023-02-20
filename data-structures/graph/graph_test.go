package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	testCases := []struct {
		vertices    []int
		numVertices int
	}{
		{
			vertices:    []int{},
			numVertices: 0,
		},
		{
			vertices:    []int{0},
			numVertices: 1,
		},
		{
			vertices:    []int{1},
			numVertices: 1,
		},
		{
			vertices:    []int{-1},
			numVertices: 1,
		},
		{
			vertices:    []int{1, 200, 300},
			numVertices: 3,
		},
		{
			vertices:    []int{1, 200, 300, 300},
			numVertices: 3,
		},
	}

	for _, tc := range testCases {
		g := Graph{}
		for _, v := range tc.vertices {
			g.AddVertex(v)
		}
		if tc.numVertices != g.numVertices {
			t.Errorf("expect '%v', but got '%v'", tc.numVertices, g.numVertices)
		}
	}
}

func TestAddEdge(t *testing.T) {
	testCases := []struct {
		vertices []int
		edges    [][]int
		adjList  map[int][]int
	}{
		{
			vertices: []int{},
			edges:    [][]int{},
		},
		{
			vertices: []int{},
			edges:    [][]int{{5, 3}},
		},
		{
			vertices: []int{5, 3},
			edges:    [][]int{},
			adjList: map[int][]int{
				5: []int{},
				3: []int{},
			},
		},
		{
			vertices: []int{5, 3},
			edges:    [][]int{{5, 0}},
			adjList: map[int][]int{
				5: []int{},
				3: []int{},
			},
		},
		{
			vertices: []int{5, 3},
			edges:    [][]int{{5, 3}},
			adjList: map[int][]int{
				3: []int{5},
				5: []int{3},
			},
		},
		{
			vertices: []int{5, 3},
			edges:    [][]int{{5, 3}, {5, 3}},
			adjList: map[int][]int{
				3: []int{5},
				5: []int{3},
			},
		},
		{
			vertices: []int{5, 3, 2},
			edges:    [][]int{{5, 3}, {5, 2}},
			adjList: map[int][]int{
				2: []int{5},
				3: []int{5},
				5: []int{3, 2},
			},
		},
		{
			vertices: []int{5, 3, 2, 4},
			edges:    [][]int{{5, 3}, {4, 2}},
			adjList: map[int][]int{
				2: []int{4},
				3: []int{5},
				4: []int{2},
				5: []int{3},
			},
		},
	}
	for _, tc := range testCases {
		g := Graph{}
		for _, v := range tc.vertices {
			g.AddVertex(v)
		}
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}
		if !reflect.DeepEqual(tc.adjList, g.adjacencyList) {
			t.Errorf("expect '%v', but got '%v'", tc.adjList, g.adjacencyList)
		}
	}
}

func TestGetNeighbors(t *testing.T) {
	type expected struct {
		v         int
		neighbors []int
	}

	testCases := []struct {
		vertices []int
		edges    [][]int
		expected []expected
	}{
		{
			vertices: []int{},
			edges:    [][]int{},
			expected: []expected{
				{2, nil},
			},
		},
		{
			vertices: []int{2, 3, 4, 5},
			edges:    [][]int{},
			expected: []expected{
				{2, []int{}},
			},
		},
		{
			vertices: []int{2, 3, 4, 5, 6, 7},
			edges: [][]int{
				{5, 3},
				{5, 4},
				{3, 2},
				{2, 4},
				{6, 7},
			},
			expected: []expected{
				{2, []int{3, 4}},
				{3, []int{5, 2}},
				{5, []int{3, 4}},
				{7, []int{6}},
			},
		},
	}

	for i, tc := range testCases {
		g := Graph{}
		for _, v := range tc.vertices {
			g.AddVertex(v)
		}
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

func TestIsAdjacent(t *testing.T) {
	type expected struct {
		uv    []int // represent vertex u and vertex v
		isAdj bool
	}

	testCases := []struct {
		vertices []int
		edges    [][]int
		expected []expected
	}{
		{
			vertices: []int{},
			edges:    [][]int{},
			expected: []expected{
				{[]int{2, 3}, false},
			},
		},
		{
			vertices: []int{2, 3, 4, 5},
			edges:    [][]int{},
			expected: []expected{
				{[]int{2, 3}, false},
			},
		},
		{
			vertices: []int{2, 3, 4, 5, 6, 7},
			edges:    [][]int{{5, 3}, {5, 4}, {3, 2}, {2, 4}, {6, 7}},
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
		g := Graph{}
		for _, v := range tc.vertices {
			g.AddVertex(v)
		}
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
