package leetcode450

import (
	"reflect"
	"testing"
)

func TestNonRecurDeleteNode(t *testing.T) {
	for _, tc := range getTestCases() {
		returnedTree := nonRecurDeleteNode(tc.originalTree, tc.val)
		if !reflect.DeepEqual(tc.updatedTree, returnedTree) {
			t.Errorf("%s - expect '%v', but got '%v'", tc.comment, tc.updatedTree, returnedTree)
		}
	}
}
