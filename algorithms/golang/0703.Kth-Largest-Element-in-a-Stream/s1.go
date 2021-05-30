package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf("top 3 is %d\n", kth.Add(3))
	fmt.Printf("top 3 is %d\n", kth.Add(5))
	fmt.Printf("top 3 is %d\n", kth.Add(10))
	fmt.Printf("top 3 is %d\n", kth.Add(9))
	fmt.Printf("top 3 is %d\n", kth.Add(4))
}

type KthLargest struct {
	k_max *IntHeap
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		k_max: &IntHeap{},
		k:     k,
	}
	for _, x := range nums {
		kth.Add(x)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.k_max.Len() < this.k {
		heap.Push(this.k_max, val)
		// fmt.Printf("1%v\n", this.k_max)
		return (*this.k_max)[0]

	}
	if (*this.k_max)[0] < val {
		(*this.k_max)[0] = val
		heap.Fix(this.k_max, 0)
		// fmt.Printf("2%v\n", this.k_max)
		return (*this.k_max)[0]

	}
	// fmt.Printf("3%v\n", this.k_max)
	return (*this.k_max)[0]
}

type IntHeap []int

// Len is the number of elements in the collection.
func (item IntHeap) Len() int {
	return len(item)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (item IntHeap) Less(i int, j int) bool {
	return item[i] < item[j]
}

// Swap swaps the elements with indexes i and j.
func (item IntHeap) Swap(i int, j int) {
	item[i], item[j] = item[j], item[i]
}

func (item *IntHeap) Push(x interface{}) {
	*item = append(*item, x.(int))
}

func (item *IntHeap) Pop() interface{} {
	old := *item
	ol := len(old)
	*item = old[0 : ol-1]
	return old[ol-1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
