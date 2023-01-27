package main

func main() {

}

// 临时数组
func rotate(nums []int, k int) {
	length := len(nums)
	// 创建等于原数组长度的切片
	temp := make([]int, length)
	// 数组复制
	copy(temp, nums)
	for i := 0; i < length; i++ {
		temp[i] = nums[i]
	}
	for i := 0; i < length; i++ {
		nums[(i+k)%length] = temp[i]
	}
}

func rotate2(nums []int, k int) {
	// 创建切片
	newNums := make([]int, len(nums))
	// i为数组下标，v为数组对应下标的元素值
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 多次反转
func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
