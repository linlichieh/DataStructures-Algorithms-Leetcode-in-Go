package leetcode139

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, s := range wordDict {
		wordMap[s] = true
	}

	// The first one is for dealing with empty string, the rest is for each char from the string
	dp := make([]bool, len(s)+1)

	// an empty string can always be segmented
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// `dp[j]` represents `s[0:j-1]` can be segmented, then find the next segment based on this foundation
			//  - index  0 1 2 3 4 5 6 7 8
			//  - string l e e t c o d e
			//  - dp     T F F F T F F F T
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	// If the last one is true, it indicates that the string is completely segmented
	return dp[len(s)]
}
