package main

func main() {

}

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	// p 为第一个大于根节点的值，即右子树的根节点
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	// 判断右子树是否都大于根节点
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}
