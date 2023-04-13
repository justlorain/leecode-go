package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for ; i < len(inorder); i++ {
		// 找到根节点在中序遍历中的位置
		if inorder[i] == preorder[0] {
			break
		}
	}
	// [ 根节点，[ 左子树的前序遍历结果 ]，[ 右子树的前序遍历结果 ] ]
	// [ [ 左子树的中序遍历结果 ]，根节点，[ 右子树的中序遍历结果 ] ]
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
