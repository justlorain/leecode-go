package main

import "fmt"

func main() {
	nums := []int{2, -1, 0, 8}
	bubbleSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
