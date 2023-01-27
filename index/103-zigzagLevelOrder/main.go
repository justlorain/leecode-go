package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var temp []int
		// 取一个临时 q
		q := queue
		// 清空 queue
		queue = nil
		for _, node := range q {
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
		res = append(res, temp)
	}
	return res
}
