package main

import "fmt"

func main() {
	fmt.Println()
}

// 自己写的
func singleNumber(nums []int) int {
	myMap := make(map[int]int)
	for _, num := range nums {
		myMap[num] += 1
	}
	for k, v := range myMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 位运算
func singleNumber2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
