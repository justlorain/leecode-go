package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	count := 0
	if nums == nil || len(nums) == 0 {
		return 0
	}
	for right := 1; right < len(nums); right++ {
		if nums[right] == nums[right-1] {
			// 出现重复的count+=1
			count++
		} else {
			nums[right-count] = nums[right]
		}
	}
	return len(nums) - count
}
