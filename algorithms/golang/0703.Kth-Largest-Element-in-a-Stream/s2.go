package main

import "sort"

type K2 struct {
	S sort.IntSlice
	k int
}

func NewK2(k int, nums []int) K2 {
	k2 := K2{
		k: k,
		S: nums,
	}

	return k2
}
