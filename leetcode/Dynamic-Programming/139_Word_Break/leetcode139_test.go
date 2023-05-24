package leetcode139

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{
			s:        "",
			wordDict: []string{""},
			expected: true,
		},
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true,
		},
	}

	for i, c := range cases {
		result := wordBreak(c.s, c.wordDict)
		if c.expected != result {
			t.Errorf("case(%d) expected %v, but got %v", i, c.expected, result)
		}
	}
}
