package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
记给定链表的长度为 nnn，注意到当向右移动的次数 k≥nk \geq nk≥n 时，我们仅需要向右移动
k % n 次即可。因为每 nnn 次移动都会让链表变为原状。
这样我们可以知道，新链表的最后一个节点为原链表的第 (n - 1) - (k % n) 个节点（从 0 开始计数）
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
