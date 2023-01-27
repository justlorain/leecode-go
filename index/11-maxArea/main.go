package main

import "math"

/*
线的数量是数组的长度
每条线的高度是数组对应下标的值

容纳的水量 = 两个指针指向的数字中较小值 * 指针之间的距离
min(x, y) * t = x * t
*/

func main() {

}

// maxArea 双指针
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ma := 0
	for l < r {
		temp := int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)
		ma = int(math.Max(float64(ma), float64(temp)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ma
}
