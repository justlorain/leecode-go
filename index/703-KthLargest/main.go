package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (k *KthLargest) Push(v any) {
	k.IntSlice = append(k.IntSlice, v.(int))
}

func (k *KthLargest) Pop() any {
	tmp := k.IntSlice[len(k.IntSlice)-1]
	k.IntSlice = k.IntSlice[:len(k.IntSlice)-1]
	return tmp
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
