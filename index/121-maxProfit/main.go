package main

import "math"

func main() {

}

// dp[i] = min(d[i-1], prices[i])
// dp[0] = prices[0]
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
