package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumWithHash(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumWithHash(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, num := range nums {
		if p, ok := hashTable[target-num]; ok {
			return []int{p, i}
		}
		hashTable[num] = i
	}
	return nil
}
