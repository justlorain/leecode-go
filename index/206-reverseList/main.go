package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 存储前一个节点
		next := curr.Next
		// 更改当前节点引用
		curr.Next = prev
		// 将当前节点设置为下一个节点的前一个节点
		prev = curr
		// 前进到下一个节点
		curr = next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
