package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 13,
		Left: &TreeNode{
			Val:   14,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
	}
	node := root
	fmt.Println(root.Left.Val)
	fmt.Println(node.Left.Val)

	node.Left.Val = -1
	fmt.Println(root.Left.Val)
	fmt.Println(node.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalWithRecursion(root *TreeNode) []int {
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

// TODO: Review
func preorderTraversalWithIteration(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val) // 根节点先加到结果集
			stack = append(stack, root) // 入栈
			root = root.Left            // 向左遍历
		}
		root = stack[len(stack)-1].Right // 取栈顶右子节点
		stack = stack[:len(stack)-1]     // 弾栈
	}
	return res
}
