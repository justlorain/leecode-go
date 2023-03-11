package main

import "fmt"

func main() {
	res := singleNumber([]int{2, 2, 3, 2})
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num] = hash[num] + 1
		} else {
			hash[num] = 1
		}
	}
	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return -1
}
