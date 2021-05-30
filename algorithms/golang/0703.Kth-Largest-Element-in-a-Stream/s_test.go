package main

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	k2 := NewK2(3, []int{6, 5, 3, 7, 8})
	sort.Sort(sort.Reverse(k2.S))
	t.Logf("%v", k2.S)
}

func TestSearch(t *testing.T) {
	k2 := NewK2(3, []int{6, 5, 3, 7, 8})
	t.Logf("%v", k2.S.Search(7))
}
