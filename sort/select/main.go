package main

func main() {

}

func sortArray(nums []int) []int {
	// 选择排序，比较交换，不稳定算法，时间O(n^2)，空间O(1)
	// 每一轮遍历，该轮的最小值前挪，从而形成前面部分是有序区
	// compare and swap
	for i := 0; i < len(nums); i++ {
		// 剪枝前面部分，比较后面部分
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
