package main

import "strings"

func main() {

}

func isPalindrome(s string) bool {
	var temp string
	for _, v := range s {
		if isAlOrNum(byte(v)) {
			temp += string(v)
		}
	}
	lower := strings.ToLower(temp)
	left, right := 0, len(lower)-1
	for left <= right {
		if lower[left] != lower[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome2(s string) bool {
	str := strings.ToLower(s)
	left, right := 0, len(str)-1
	for left < right {
		for left < right && !isAlOrNum(str[left]) {
			left++
		}
		for left < right && !isAlOrNum(str[right]) {
			right--
		}
		if left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isAlOrNum(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
