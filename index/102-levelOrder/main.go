package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		s = append(s, 4)
		fmt.Println(i)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderWithBFS(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root} // 根节点入队
	idx := 0
	for len(queue) > 0 { // 队列不为空就一直循环
		res = append(res, []int{})
		var temp []*TreeNode
		for j := 0; j < len(queue); j++ { // 遍历队列
			node := queue[j]
			res[idx] = append(res[idx], node.Val) // 队列中的元素为同一层的节点
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		queue = temp // 切换到下一层的队列
		idx++
	}
	return res
}
