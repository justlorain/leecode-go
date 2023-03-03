package main

import "math"

func main() {

}

// dp[i] = min(d[i-1], prices[i])
// dp[0] = prices[0]
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	// 一次遍历：寻找股票最低点，然后计算当天利润
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}
