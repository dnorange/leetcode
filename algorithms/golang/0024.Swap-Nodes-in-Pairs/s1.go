package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%d=>%s", l.Val, l.Next)
}

func swapPairs(head *ListNode) *ListNode {
	self := &ListNode{
		Next: head,
	}
	current := self
	for current.Next != nil && current.Next.Next != nil {
		a := current.Next
		b := a.Next

		a.Next = b.Next
		b.Next = a
		current.Next = b

		current = a
	}
	return self.Next
}
