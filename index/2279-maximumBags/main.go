package main

import "sort"

func main() {

}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	res := 0
	// 先计算每个背包还可以装多少石头
	for i := range capacity {
		capacity[i] -= rocks[i]
	}
	// 对背包进行排序
	sort.Ints(capacity)
	// 从小到大装石头
	for _, leftSpace := range capacity {
		if leftSpace > additionalRocks {
			// 无法装满，说明后序所有的都无法装满，因为已经进行了排序
			break
		}
		res++
		additionalRocks -= leftSpace
	}
	return res
}
