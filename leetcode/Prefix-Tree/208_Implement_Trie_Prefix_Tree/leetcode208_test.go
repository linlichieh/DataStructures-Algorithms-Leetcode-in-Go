package leetcode208

import (
	"reflect"
	"testing"
)

func TestXXX(t *testing.T) {
	cases := []struct {
		word      string
		prefix    string
		expected1 bool
		expected2 bool
	}{
		{
			word:      "apple",
			prefix:    "apple",
			expected1: true,
			expected2: true,
		},
		{
			word:      "apple",
			prefix:    "app",
			expected1: true,
			expected2: true,
		},
		{
			word:      "apple",
			prefix:    "appel",
			expected1: true,
			expected2: false,
		},
		{
			word:      "apple",
			prefix:    "appe",
			expected1: true,
			expected2: false,
		},
		{
			word:      "apple",
			prefix:    "orange",
			expected1: true,
			expected2: false,
		},
	}

	for i, c := range cases {
		obj := Constructor()
		obj.Insert(c.word)
		result1 := obj.Search(c.word)
		result2 := obj.StartsWith(c.prefix)
		if !reflect.DeepEqual(c.expected1, result1) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected1, result1)
		}
		if !reflect.DeepEqual(c.expected2, result2) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected2, result2)
		}
	}
}
