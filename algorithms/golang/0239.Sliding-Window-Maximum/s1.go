package main

import "container/heap"

// timeout solution with pq
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)

	vals := make([]int, k)
	copy(vals, nums[:k])
	mySlice := MyIntSlice(vals)
	heap.Init(&mySlice)

	result = append(result, mySlice[0])

	for i := k; i < len(nums); i++ {
		//oriRemoveIndex++
		realRange := nums[i-k+1 : i+1]
		mySlice.Push(nums[i])
		heap.Fix(&mySlice, mySlice.Len()-1)
		//fmt.Printf("m%v r%v k%v i%v ", mySlice, realRange, k, i)
		result = append(result, pop(&mySlice, realRange))
		//fmt.Printf("\n")
	}
	return result
}

func pop(mySlice *MyIntSlice, realRange []int) int {
	for _, x := range realRange {
		// fmt.Printf("s%v r%v ", mySlice, realRange)
		if (*mySlice)[0] == x {
			return (*mySlice)[0]
		}
	}
	heap.Remove(mySlice, 0)
	//fmt.Printf("p%v", mySlice.Pop())

	return pop(mySlice, realRange)
}

type MyIntSlice []int

func (s *MyIntSlice) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *MyIntSlice) Pop() interface{} {
	old := *s
	oldLen := len(old)
	*s = old[:oldLen-1]
	return old[oldLen-1]
}

func (s MyIntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s MyIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MyIntSlice) Len() int {
	return len(s)
}
