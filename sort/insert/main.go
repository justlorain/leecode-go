package main

func main() {

}

func sortArray(nums []int) []int {
	// 插入排序，比较交换，稳定算法，时间O(n^2)，空间O(1)
	// 0->len方向，每轮从后往前比较，相当于找到合适位置，插入进去
	// 数据规模小的时候，或基本有序，效率高
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp { //左边比右边大
			nums[j+1] = nums[j] //右移1位
			j--                 //扫描前一个数
		}
		nums[j+1] = tmp //添加到小于它的数的右边
	}
	return nums
}
