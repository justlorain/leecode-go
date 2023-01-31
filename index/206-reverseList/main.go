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

func reverseListWithRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
