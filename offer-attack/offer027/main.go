package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	sli := make([]int, 0)
	for head != nil {
		sli = append(sli, head.Val)
		head = head.Next
	}
	left, right := 0, len(sli)-1
	for left <= right {
		if sli[left] != sli[right] {
			return false
		}
		left++
		right--
	}
	return true
}
