package main

import "fmt"

func main() {
	// 4
	// 1 2 3 4
	var l int
	var arr []int
	_, _ = fmt.Scanln(&l)
	for i := 0; i < l; i++ {
		var tmp int
		_, _ = fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	fmt.Println(l, arr)
}
