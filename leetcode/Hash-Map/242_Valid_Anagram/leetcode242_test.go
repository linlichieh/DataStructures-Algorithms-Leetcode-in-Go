package leetcode242

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "",
			t:        "",
			expected: true,
		},
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			s:        "你好嗎?",
			t:        "好?嗎你",
			expected: true,
		},
		{
			s:        "你好嗎?",
			t:        "好嗎!你",
			expected: false,
		},
		{
			s:        "你好好嗎?",
			t:        "好嗎?你好",
			expected: true,
		},
		{
			s:        "你好好嗎?",
			t:        "好嗎!你好",
			expected: false,
		},
		{
			s:        "你好好嗎?",
			t:        "好嗎!你",
			expected: false,
		},
	}

	for i, c := range cases {
		result := isAnagram(c.s, c.t)
		if result != c.expected {
			t.Errorf("(%d) expected: '%v', but got '%v'", i, c.expected, result)
		}
	}
}
