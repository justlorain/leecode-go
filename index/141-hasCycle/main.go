package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleWithHash(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycleWithFastSlowPointer(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
