package leetcode49

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		input  []string
		output [][]string
	}{
		{
			input:  []string{},
			output: [][]string{},
		},
		{
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			input:  []string{""},
			output: [][]string{{""}},
		},
		{
			input:  []string{"a"},
			output: [][]string{{"a"}},
		},
	}

	for i, c := range cases {
		result := groupAnagrams(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
