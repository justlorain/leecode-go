package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		res = pathP[i]
	}
	return res
}

func getPath(root, target *TreeNode) []*TreeNode {
	node := root
	path := make([]*TreeNode, 0)
	for node != target {
		path = append(path, node)
		if target.Val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	path = append(path, node)
	return path
}
