package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenWithRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	var preorder func(root *TreeNode) []*TreeNode
	preorder = func(root *TreeNode) []*TreeNode {
		var list []*TreeNode
		if root != nil {
			list = append(list, root)
			list = append(list, preorder(root.Left)...)
			list = append(list, preorder(root.Right)...)
		}
		return list
	}
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func flattenWithIteration(root *TreeNode) {
	var list []*TreeNode
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}
