package main

func main() {

}

func minMoves(nums []int) int {
	min := nums[0]
	res := 0
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		res += num - min
	}
	return res
}
