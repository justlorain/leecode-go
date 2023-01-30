package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	hash := make(map[byte]int)
	length := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < length; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(hash, s[i-1])
		}
		for rk+1 < length && hash[s[rk+1]] == 0 {
			// 不断地移动右指针
			hash[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
