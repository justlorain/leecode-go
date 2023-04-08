package main

import "fmt"

func main() {
	sli := []int{-1, 4, 2, 3}
	fmt.Println(checkPossibility(sli))

	hash := map[int]int{
		1: 2,
		2: 3,
	}
	fmt.Println(hash[1])
}

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
