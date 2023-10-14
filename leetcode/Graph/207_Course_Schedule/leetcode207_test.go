package leetcode207

import "testing"

func TestCanFinish(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{2, 1}, {1, 0}},
			expected:      true,
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}},
			expected:      true,
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			expected:      false,
		},
		{
			numCourses:    5,
			prerequisites: [][]int{{1, 2}, {2, 0}, {3, 1}, {4, 3}},
			expected:      true,
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{0, 1}, {2, 3}},
			expected:      true,
		},
	}

	for i, c := range cases {
		result := canFinish(c.numCourses, c.prerequisites)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
