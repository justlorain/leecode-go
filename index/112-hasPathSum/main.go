package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSumWithDFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSumWithDFS(root.Left, targetSum-root.Val) || hasPathSumWithDFS(root.Right, targetSum-root.Val)
}

func hasPathSumWithBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 根节点入队
	qNode := []*TreeNode{root}
	qVal := []int{root.Val}
	for len(qNode) != 0 {
		// 保存当前节点和当前节点值
		now := qNode[0]
		nowVal := qVal[0]
		// 出队
		qNode = qNode[1:]
		qVal = qVal[1:]
		if now.Left == nil && now.Right == nil {
			if nowVal == targetSum {
				return true
			}
			continue
		}
		if now.Left != nil {
			qNode = append(qNode, now.Left)
			qVal = append(qVal, nowVal+now.Left.Val)
		}
		if now.Right != nil {
			qNode = append(qNode, now.Right)
			qVal = append(qVal, nowVal+now.Right.Val)
		}
	}
	return false
}
