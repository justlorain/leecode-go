package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*TreeNode)
		ans = append(ans, tmp.Val)
		queue.Remove(queue.Front())
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return ans
}

func levelOrder2(root *TreeNode) []int {
	if root == nil {
		return []int{0}
	}
	var ans []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		ans = append(ans, node.Val)
		queue.Remove(queue.Front())
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ans
}
