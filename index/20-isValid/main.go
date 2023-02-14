package main

import "fmt"

func main() {
	//for _, v := range s {
	//	fmt.Println(v)
	//	fmt.Printf("%T", v) // int32
	//}
	s := "()"
	pairs := map[byte]byte{
		'(': '(',
		']': '[',
		'}': '{',
	}
	fmt.Println(pairs[s[0]])

	//fmt.Println(s[1])  // 101
	//fmt.Println(s[1:]) // ello
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if _, ok := pairs[s[i]]; ok {
			//  栈为空 || 括号不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}