package main

func main() {

}

func missingNumber(nums []int) int {
	hash := make(map[int]bool)
	for _, v := range nums {
		hash[v] = true
	}
	for i := 0; ; i++ {
		if !hash[i] {
			return i
		}
	}
}

func missingNumber2(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}
