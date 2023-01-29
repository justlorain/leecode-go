package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalWithRecursion(root *TreeNode) []int {
	var res []int
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

func postorderTraversalWithRecursion2(root *TreeNode) []int {
	var postorder func(node *TreeNode) []int
	postorder = func(node *TreeNode) []int {
		var res []int
		if node == nil {
			return res
		}
		res = append(res, postorder(node.Left)...)
		res = append(res, postorder(node.Right)...)
		res = append(res, node.Val)
		return res
	}
	return postorder(root)
}

// TODO: Review
func postorderTraversalWithIteration(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   // 赋值为栈顶
		stack = stack[:len(stack)-1] // 弾栈
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}
