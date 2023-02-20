package stack

import (
	"testing"
)

func TestLinkedListStackPush(t *testing.T) {
	testCases := []struct {
		input []string
		top   string
		size  int
	}{
		{
			input: []string{},
			top:   "",
			size:  0,
		},
		{
			input: []string{"google"},
			top:   "google",
			size:  1,
		},
		{
			input: []string{"google", "microsoft"},
			top:   "microsoft",
			size:  2,
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			top:   "amazon",
			size:  3,
		},
	}

	for _, tc := range testCases {
		stack := NewStack("linked_list")
		for _, val := range tc.input {
			stack.Push(val)
		}
		if tc.size != stack.Size() {
			t.Errorf("Size: expect '%v', but got '%v'", tc.size, stack.Size())
		}
		if tc.top != stack.Top() {
			t.Errorf("Top: expect '%v', but got '%v'", tc.top, stack.Top())
		}
	}
}

func TestLinkedListStackPop(t *testing.T) {
	type result struct {
		pop  string
		top  string
		size int
	}

	testCases := []struct {
		input  []string
		result []result
	}{
		{
			input: []string{},
			result: []result{
				{pop: "", top: "", size: 0},
				{pop: "", top: "", size: 0},
			},
		},
		{
			input: []string{"google"},
			result: []result{
				{pop: "google", top: "", size: 0},
				{pop: "", top: "", size: 0},
			},
		},
		{
			input: []string{"google", "microsoft"},
			result: []result{
				{pop: "microsoft", top: "google", size: 1},
				{pop: "google", top: "", size: 0},
			},
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			result: []result{
				{pop: "amazon", top: "microsoft", size: 2},
				{pop: "microsoft", top: "google", size: 1},
				{pop: "google", top: "", size: 0},
			},
		},
	}

	for _, tc := range testCases {
		stack := NewStack("linked_list")
		for _, val := range tc.input {
			stack.Push(val)
		}
		for i, expected := range tc.result {
			// pop
			popVal := stack.Pop()
			if expected.pop != popVal {
				t.Errorf("Pop(%d): expect '%v', but got '%v'", i, expected.pop, popVal)
			}

			// size
			if expected.size != stack.Size() {
				t.Errorf("Size(%d): expect '%v', but got '%v'", i, expected.size, stack.Size())
			}

			// top
			if expected.top != stack.Top() {
				t.Errorf("Pop(%d) expect '%v', but got '%v'", i, expected.top, stack.Top())
			}
		}
	}
}

func TestLinkedListStackTop(t *testing.T) {
	testCases := []struct {
		input []string
		top   string
	}{
		{
			input: []string{},
			top:   "",
		},
		{
			input: []string{"google"},
			top:   "google",
		},
		{
			input: []string{"google", "microsoft"},
			top:   "microsoft",
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			top:   "amazon",
		},
	}
	for _, tc := range testCases {
		stack := NewStack("linked_list")
		for _, val := range tc.input {
			stack.Push(val)
		}
		if tc.top != stack.Top() {
			t.Errorf("expect '%v', but got '%v'", tc.top, stack.Top())
		}
	}
}
