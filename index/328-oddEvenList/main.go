package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 偶数链表的头结点
	evenHead := head.Next
	// 奇数节点
	odd := head
	// 偶数节点
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
