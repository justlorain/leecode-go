package main

import "sort"

func main() {

}

// [0, 1, 2, 4]
func missingNumber(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func missingNumberWithHash(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	for i := 0; ; i++ {
		if !set[i] {
			return i
		}
	}
}
