package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var inorder func(node *TreeNode)
	res := 0
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		if k == 0 {
			return
		}
		res = node.Val
		k--
		inorder(node.Left)
	}
	inorder(root)
	return res
}
