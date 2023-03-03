package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 2, 3, 1}
	sorted := sortArrayWithMergeSort(nums)
	fmt.Println(sorted)
}

func sortArrayWithQuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	// 边界条件，若l > r则非法，若l=r则只有一个元素而无需处理
	if left >= right {
		return
	}

	// 找[right, left]的随机数
	randIdx := left + rand.Intn(right-left+1)
	nums[left], nums[randIdx] = nums[randIdx], nums[left] // 交换: 提升接近有序的case性能
	pivot := nums[left]

	// partition的过程，让a[:idx]<a[0] && a[idx:]>=nums[0]
	idx := left
	// all in nums[left+1, idx] <= pivot
	// all in nums(idx, right] > pivot
	for i := left + 1; i <= right; i++ { // i从left+1开始，是因为最终要将nums[left]和nums[idx]交换
		if nums[i] <= pivot { // 大不管，小放前，等号无所谓
			idx++
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	// partition后，将a[left]放在正确的位置
	nums[left], nums[idx] = nums[idx], nums[left]

	// 规模变小并分治
	quickSort(nums, left, idx-1)
	quickSort(nums, idx+1, right)
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
		// merge
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
		copy(nums[left:right+1], temp)
		//for idx := 0; idx < right-left+1; idx++ {
		//	nums[idx+left] = temp[idx]
		//}
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

// prefer this one
func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i, j := left, mid+1
	var tmp []int
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	if i <= mid {
		tmp = append(tmp, nums[i:mid+1]...)
	} else {
		tmp = append(tmp, nums[j:right+1]...)
	}
	copy(nums[left:right+1], tmp)
}

func sortArrayWithMergeSort3(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		sortArrayWithMergeSort3(nums, left, mid)
		sortArrayWithMergeSort3(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge3(nums []int, left, mid, right int) {
	i, j := left, mid+1
	var tmp []int
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	if i <= mid {
		tmp = append(tmp, nums[i:mid+1]...)
	} else {
		tmp = append(tmp, nums[j:right+1]...)
	}
	copy(nums[left:right+1], tmp)
}
