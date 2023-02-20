package main

import "fmt"

func main() {
	fmt.Println(mySqrt(15))
}

func mySqrt(x int) int {
	left, right := 0, x
	res := -1
	for left <= right {
		// mid = (low + high) / 2
		// mid = low + (high - low) / 2
		// mid = low + ((high - low) >> 1)
		mid := left + (right-left)/2
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
