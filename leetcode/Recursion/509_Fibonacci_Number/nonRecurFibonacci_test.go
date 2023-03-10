package leetcode509

import "testing"

func TestNonRecurFib(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		result := nonRecurFib(tc.input)
		if result != tc.output {
			t.Errorf("expect '%v', but got '%v'", tc.output, result)
		}
	}
}
