package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range numbers {
		if v, ok := hash[target-num]; ok {
			return []int{v, i}
		}
		hash[num] = i
	}
	return nil
}
