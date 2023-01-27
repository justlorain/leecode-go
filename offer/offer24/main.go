package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // 暂存后继节点
		curr.Next = prev  // 修改当前节点引用指向
		prev = curr       // 暂存当前节点
		curr = next       // 访问下一节点
	}
	return prev
}
