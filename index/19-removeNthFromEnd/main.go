package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	// 这里不直接返回 head 的是为了应对链表只有 head 一个元素并且删除倒数第 1 个元素那么 head 就会为空
	return dummy.Next
}
