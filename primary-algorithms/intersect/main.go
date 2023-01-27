package main

import "sort"

func main() {

}

// 排序+双指针
func intersect(nums1 []int, nums2 []int) []int {
	var temp []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	left := 0
	right := 0
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] < nums2[right] {
			left++
		} else if nums1[left] > nums2[right] {
			right++
		} else {
			temp = append(temp, nums1[left])
			left++
			right++
		}
	}
	return temp
}

// 哈希表
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}
