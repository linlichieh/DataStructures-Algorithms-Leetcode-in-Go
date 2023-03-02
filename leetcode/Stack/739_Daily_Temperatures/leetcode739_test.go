package leetcode739

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:  []int{30, 40, 50, 60},
			output: []int{1, 1, 1, 0},
		},
		{
			input:  []int{30, 60, 90},
			output: []int{1, 1, 0},
		},
	}

	for _, c := range cases {
		result := dailyTemperatures(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
