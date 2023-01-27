package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 栈顶元素
	top := this.minStack[len(this.minStack)-1]
	// 如果新入栈的元素比栈顶元素大则继续append栈顶元素
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	// 利用切片切到最后一个元素之前，也就是最后append到数组中的元素
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
