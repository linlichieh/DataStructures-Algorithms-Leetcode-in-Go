package queue

import "testing"

func TestLinkedListQueueEnqueue(t *testing.T) {
	testCases := []struct {
		input []string
		peek  string
		size  int
	}{
		{
			input: []string{},
			peek:  "",
			size:  0,
		},
		{
			input: []string{"google"},
			peek:  "google",
			size:  1,
		},
		{
			input: []string{"google", "microsoft"},
			peek:  "google",
			size:  2,
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			peek:  "google",
			size:  3,
		},
	}

	for _, tc := range testCases {
		queue := NewQueue("linked_list")
		for _, val := range tc.input {
			queue.Enqueue(val)
		}
		if tc.size != queue.Size() {
			t.Errorf("Size: expect '%v', but got '%v'", tc.size, queue.Size())
		}
		if tc.peek != queue.Peek() {
			t.Errorf("Top: expect '%v', but got '%v'", tc.peek, queue.Peek())
		}
	}
}

func TestLinkedListQueueDequeue(t *testing.T) {
	type result struct {
		dequeue string
		peek    string
		size    int
	}

	testCases := []struct {
		input  []string
		result []result
	}{
		{
			input: []string{},
			result: []result{
				{dequeue: "", peek: "", size: 0},
				{dequeue: "", peek: "", size: 0},
			},
		},
		{
			input: []string{"google"},
			result: []result{
				{dequeue: "google", peek: "", size: 0},
				{dequeue: "", peek: "", size: 0},
			},
		},
		{
			input: []string{"google", "microsoft"},
			result: []result{
				{dequeue: "google", peek: "microsoft", size: 1},
				{dequeue: "microsoft", peek: "", size: 0},
			},
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			result: []result{
				{dequeue: "google", peek: "microsoft", size: 2},
				{dequeue: "microsoft", peek: "amazon", size: 1},
				{dequeue: "amazon", peek: "", size: 0},
			},
		},
	}

	for _, tc := range testCases {
		queue := NewQueue("linked_list")
		for _, val := range tc.input {
			queue.Enqueue(val)
		}
		for i, expected := range tc.result {
			// dequeue
			val := queue.Dequeue()
			if expected.dequeue != val {
				t.Errorf("Dequeue(%d): expect '%v', but got '%v'", i, expected.dequeue, val)
			}

			// size
			if expected.size != queue.Size() {
				t.Errorf("Size(%d): expect '%v', but got '%v'", i, expected.size, queue.Size())
			}

			// peek
			if expected.peek != queue.Peek() {
				t.Errorf("Peek(%d) expect '%v', but got '%v'", i, expected.peek, queue.Peek())
			}
		}
	}
}

func TestLinkedListQueuePeek(t *testing.T) {
	testCases := []struct {
		input []string
		peek  string
	}{
		{
			input: []string{},
			peek:  "",
		},
		{
			input: []string{"google"},
			peek:  "google",
		},
		{
			input: []string{"google", "microsoft"},
			peek:  "google",
		},
		{
			input: []string{"google", "microsoft", "amazon"},
			peek:  "google",
		},
	}
	for _, tc := range testCases {
		queue := NewQueue("linked_list")
		for _, val := range tc.input {
			queue.Enqueue(val)
		}
		if tc.peek != queue.Peek() {
			t.Errorf("expect '%v', but got '%v'", tc.peek, queue.Peek())
		}
	}
}
