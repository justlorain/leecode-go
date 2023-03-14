package main

func main() {

}

func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + ((right - left) >> 1)
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i, j := left, mid+1
	temp := make([]int, 0)
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	if i <= mid {
		temp = append(temp, nums[i:mid+1]...)
	} else {
		temp = append(temp, nums[j:right+1]...)
	}
	copy(nums[left:right+1], temp)
}
