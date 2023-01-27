package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	// 用于节点的出入栈
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// 长度为多少即该层有多少元素
		l := len(queue)
		var levelList []int
		for i := 0; i < l; i++ {
			node := queue[0]
			// 值加入levelList
			levelList = append(levelList, node.Val)
			// 删除
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//该层的值数组加入res
		res = append(res, levelList)
	}
	return res
}
