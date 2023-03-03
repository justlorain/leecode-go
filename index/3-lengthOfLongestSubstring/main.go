package main

import (
	"fmt"
	"math"
)

func main() {
	//s := "hello"
	//for _, v := range s {
	//	fmt.Printf("%T\n", v) // int32 (rune)
	//	fmt.Printf("%v\n", v) // 104
	//	break
	//}
	//fmt.Printf("%T\n", s[0:3]) // string
	//fmt.Printf("%v\n", s[0:3]) // hel
	//
	//fmt.Printf("%T\n", s[0]) // uint8 (byte)
	//fmt.Printf("%v\n", s[0]) // 104

	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring3(s))

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

func lengthOfLongestSubstring2(s string) int {
	hash := make(map[byte]int)
	length := len(s)
	right, res := -1, 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(hash, s[i-1])
		}
		for right+1 < length && hash[s[right+1]] == 0 {
			hash[s[right+1]]++
			right++
		}
		res = int(math.Max(float64(res), float64(right-i+1)))
	}
	return res
}

// TODO: review
func lengthOfLongestSubstring3(s string) int {
	res, left := 0, 0
	cnt := make(map[byte]int)
	for right, c := range s {
		cnt[byte(c)]++
		for cnt[byte(c)] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
