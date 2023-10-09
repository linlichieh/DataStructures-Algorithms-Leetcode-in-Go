package leetcode121

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	var profit int
	for _, price := range prices {
		profit = price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}
