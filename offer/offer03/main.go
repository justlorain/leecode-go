package main

func findRepeatNumber(nums []int) int {
	set := make(map[int]int)
	for _, num := range nums {
		if _, has := set[num]; has {
			return num
		}
		set[num]++
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return num
		}
		set[num] = struct{}{} // struct{}表示是结构体类型，后面的{}是具体的内容
	}
	return -1
}
