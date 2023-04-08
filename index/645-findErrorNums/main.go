package main

import "sort"

func main() {

}

func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	sort.Ints(nums)
	pre := 0
	for _, num := range nums {
		if num == pre {
			res[0] = num
		} else if num-pre > 1 {
			res[1] = pre + 1
		}
		pre = num
	}
	if nums[len(nums)-1] != len(nums) {
		res[1] = len(nums)
	}
	return res
}

func findErrorNumsWithHash(nums []int) []int {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	res := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := cnt[i]; c == 2 {
			res[0] = i
		} else if c == 0 {
			res[1] = i
		}
	}
	return res
}
