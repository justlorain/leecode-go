package contains_duplicate

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
}

// 排序，然后判断相邻位置的元素是否相等
func containsDuplicate(nums []int) bool {
	sort.Ints(nums) // 排序
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 哈希表
func containsDuplicate2(nums []int) bool {
	//set := map[int]struct{}{}
	set := make(map[int]struct{})
	for _, v := range nums {
		// 这里先是对has赋值，如果存在就是true然后再将true也就是has作为if语句的判断条件使用
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
