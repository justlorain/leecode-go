package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 1}
	sorted := sortArrayWithMergeSort(nums)
	fmt.Println(sorted)
}

func sortArrayWithMergeSort(nums []int) []int {
	temp := make([]int, len(nums))
	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		// mid = low + ((high - low) >> 1)
		mid := left + ((right - left) >> 1)
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		i, j, k := left, mid+1, 0
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				temp[k] = nums[i]
				k++
				i++
			} else {
				temp[k] = nums[j]
				k++
				j++
			}
		}
		for i <= mid {
			temp[k] = nums[i]
			k++
			i++
		}
		for j <= right {
			temp[k] = nums[j]
			k++
			j++
		}
		for idx := 0; idx < right-left+1; idx++ {
			nums[idx+left] = temp[idx]
		}
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
