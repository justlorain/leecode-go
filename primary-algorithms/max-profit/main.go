package main

func main() {

}

func maxProfit(prices []int) int {
	length := len(prices)
	sum := 0
	if len(prices) == 0 {
		return 0
	}
	for i := 0; i < length-1; i++ {
		count := prices[i+1] - prices[i]
		if count > 0 {
			sum += count
		}
	}
	return sum
}
