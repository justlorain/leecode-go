package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[:len(arr)-1])
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalWithRecursion(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// TODO: Review
func inorderTraversalWithIteration(root *TreeNode) (res []int) {
	var stack []*TreeNode
	// 有一个 true 就是 true，全为 false 才为 false
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   // root 赋值为 栈顶元素
		stack = stack[:len(stack)-1] // 弾栈
		res = append(res, root.Val)  // 将弹出的元素添加到结果集中
		root = root.Right            // root 赋值为右子节点
	}
	return
}
