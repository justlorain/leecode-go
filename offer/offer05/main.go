package main

import (
	"fmt"
	"strings"
)

func main() {
	s := replaceSpace3("We are happy.")
	fmt.Println(s)

	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a[:2])

}

func replaceSpace(s string) string {
	var res string
	for _, v := range s {
		switch v {
		case ' ':
			res += "%20"
		default:
			res += string(v)
		}
	}
	return res
}

func replaceSpace2(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func replaceSpace3(s string) string {
	sb := strings.Builder{}
	for _, v := range s {
		switch v {
		case ' ':
			sb.WriteString("%20")
		default:
			sb.WriteString(string(v))
		}
	}
	return sb.String()
}
