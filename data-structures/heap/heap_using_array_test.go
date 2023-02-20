package heap

import (
	"reflect"
	"testing"
)

type minHeapInsertTestCase struct {
	input    []int
	expected []int
}

type minHeapExtractMinTestCase struct {
	input    []int
	expected [][]int // Each result corresponds to each ExtractMin
}

func getMinHeapInsertTestCases() []minHeapInsertTestCase {
	return []minHeapInsertTestCase{
		{
			input: []int{},
			// make expected as nil so that it can pass the test
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{1, 3, 2},
		},
		{
			input:    []int{3, 1, 2},
			expected: []int{1, 3, 2},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			input:    []int{2, 1, 2},
			expected: []int{1, 2, 2},
		},
		{
			input:    []int{-2, -1, -2},
			expected: []int{-2, -1, -2},
		},
		{
			input:    []int{-2, -3, -1},
			expected: []int{-3, -2, -1},
		},
		{
			input:    []int{34, 21, 55, 23, 88, 2, 66, 35, 1, 77},
			expected: []int{1, 2, 21, 23, 77, 55, 66, 35, 34, 88},
		},
		{
			input:    []int{34, -21, 55, 23, -88, 2, 66, -35, 1, 77},
			expected: []int{-88, -35, 2, -21, 23, 55, 66, 34, 1, 77},
		},
	}
}

func getMinHeapExtractMinTestCases() []minHeapExtractMinTestCase {
	cases := []minHeapExtractMinTestCase{
		{
			input: []int{},
			expected: [][]int{
				nil,
			},
		},
		{
			input: []int{1},
			expected: [][]int{
				[]int{},
			},
		},
		{
			input: []int{1, 2},
			expected: [][]int{
				[]int{2},
				[]int{},
			},
		},
		{
			input: []int{3, 1, 2},
			expected: [][]int{
				[]int{2, 3},
				[]int{3},
				[]int{},
			},
		},
		{
			input: []int{2, 2, 2},
			expected: [][]int{
				[]int{2, 2},
				[]int{2},
				[]int{},
			},
		},
		{
			input: []int{-2, -1, -2},
			expected: [][]int{
				[]int{-2, -1},
				[]int{-1},
				[]int{},
			},
		},
		{
			input: []int{-2, -3, -1},
			expected: [][]int{
				[]int{-2, -1},
				[]int{-1},
				[]int{},
			},
		},
		{
			input: []int{34, 21, 55, 23, 88, 2, 66, 35, 1, 77}, // after sort: [1, 2, 21, 23, 77, 55, 66, 35, 34, 88]
			expected: [][]int{
				[]int{2, 23, 21, 34, 77, 55, 66, 35, 88},
				[]int{21, 23, 55, 34, 77, 88, 66, 35},
				[]int{23, 34, 55, 35, 77, 88, 66},
				[]int{34, 35, 55, 66, 77, 88},
				[]int{35, 66, 55, 88, 77},
				[]int{55, 66, 77, 88},
				[]int{66, 88, 77},
				[]int{77, 88},
				[]int{88},
				[]int{},
			},
		},
	}
	return []minHeapExtractMinTestCase{cases[7]}
}

func TestInsert(t *testing.T) {
	cases := getMinHeapInsertTestCases()

	for i, c := range cases {
		heap := MinHeapArray{}
		for _, v := range c.input {
			heap.Insert(v)
		}
		if !reflect.DeepEqual(c.expected, heap.list) {
			t.Errorf("(%d) expected '%v', but got '%v'", i, c.expected, heap.list)
		}
	}
}

func TestExtractTest(t *testing.T) {
	cases := getMinHeapExtractMinTestCases()

	for i, c := range cases {
		heap := MinHeapArray{}
		for _, v := range c.input {
			heap.Insert(v)
		}
		for _, expected := range c.expected {
			heap.ExtractMin()
			if !reflect.DeepEqual(expected, heap.list) {
				t.Errorf("(%d) expected '%v', but got '%v'", i, expected, heap.list)
			}
		}
	}
}
