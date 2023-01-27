package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	// 空结构体不占据空间
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		// ensure that set1 is the smaller one
		set1, set2 = set2, set1
	}
	for k := range set1 {
		if _, has := set2[k]; has {
			ans = append(ans, k)
		}
	}
	return ans
}
