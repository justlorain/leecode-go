package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	idx := 0
	for len(queue) > 0 {
		var temp []*TreeNode
		// 这里一定要初始化一下
		res = append(res, []int{})
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			res[idx] = append(res[idx], node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		queue = temp
		idx++
	}
	// 反转数组
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	// res
	return res
}
