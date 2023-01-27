package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	// 反转数组
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func reversePrint2(head *ListNode) []int {
	var prev *ListNode
	curr := head
	var ans []int
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	temp := prev
	for temp != nil {
		ans = append(ans, temp.Val)
		temp = temp.Next
	}
	return ans
}
