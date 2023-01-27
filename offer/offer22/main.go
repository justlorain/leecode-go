package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
链表的尾结点的 Next 为 nil
*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
