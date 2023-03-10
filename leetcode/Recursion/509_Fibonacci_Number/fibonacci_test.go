package leetcode509

import "testing"

type testCase struct {
	input  int
	output int
}

func getTestCases() []testCase {
	return []testCase{
		{
			input:  0,
			output: 0,
		},
		{
			input:  1,
			output: 1,
		},
		{
			input:  2,
			output: 1,
		},
		{
			input:  3,
			output: 2,
		},
		{
			input:  4,
			output: 3,
		},
		{
			input:  5,
			output: 5,
		},
		{
			input:  8,
			output: 21,
		},
		{
			input:  10,
			output: 55,
		},
		{
			input:  15,
			output: 610,
		},
	}
}

func TestFib(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		result := fib(tc.input)
		if result != tc.output {
			t.Errorf("expect '%v', but got '%v'", tc.output, result)
		}
	}
}
