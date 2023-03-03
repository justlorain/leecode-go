package main

func main() {

}

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num] += 1
	}
	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return -1
}
