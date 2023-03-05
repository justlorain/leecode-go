package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hash := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycleWithFastSlowPointer(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
