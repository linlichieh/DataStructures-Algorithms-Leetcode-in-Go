package leetcode98

import "testing"

func TestIsValid2(t *testing.T) {
	cases := getTestCases()
	for i, c := range cases {
		isValid := isValidBST2(c.tree)
		if isValid != c.isValid {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.isValid, isValid)
		}
	}
}
