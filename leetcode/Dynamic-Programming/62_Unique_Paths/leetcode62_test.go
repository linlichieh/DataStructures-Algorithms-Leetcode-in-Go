package leetcode62

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m      int // row
		n      int // col
		output int
	}{
		{
			m:      0,
			n:      0,
			output: 0,
		},
		{
			m:      3,
			n:      1,
			output: 1,
		},
		{
			m:      1,
			n:      2,
			output: 1,
		},
		{
			m:      3,
			n:      2,
			output: 3,
		},
		{
			m:      4,
			n:      2,
			output: 4,
		},
		{
			m:      3,
			n:      3,
			output: 6,
		},
		{
			m:      4,
			n:      4,
			output: 20,
		},
		{
			m:      4,
			n:      3,
			output: 10,
		},
		{
			m:      5,
			n:      2,
			output: 5,
		},
		{
			m:      3,
			n:      7,
			output: 28,
		},
		{
			m:      51,
			n:      9,
			output: 1916797311,
		},
	}

	for i, c := range cases {
		result := uniquePaths(c.m, c.n)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
