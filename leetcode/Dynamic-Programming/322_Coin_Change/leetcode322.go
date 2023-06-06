package leetcode322

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)

	// Set default values
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// Iterate the dp
	for i := 1; i <= amount; i++ {
		// Loop the coins
		for _, coin := range coins {
			// Build the cache from bottom by subtracting coin from index i
			if i-coin >= 0 {
				// Update the dp with the minimum value
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	// If the index `amount` is still the same as default value, which is amount +1, it means that it can't be
	if dp[amount] != amount+1 {
		return dp[amount]
	}

	return -1
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
