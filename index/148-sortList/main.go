package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// 寻找链表的中点使用快慢指针：
	// 快指针每次移动两步，慢指针每次移动一步，快指针到达末尾时慢指针指向的就是中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(h1, h2 *ListNode) *ListNode {
	temp := &ListNode{}
	p := temp
	for h1 != nil && h2 != nil {
		if h1.Val >= h2.Val {
			p.Next = h2
			h2 = h2.Next
		} else {
			p.Next = h1
			h1 = h1.Next
		}
		p = p.Next
	}
	if h1 != nil {
		p.Next = h1
	}
	if h2 != nil {
		p.Next = h2
	}
	return temp.Next
}
