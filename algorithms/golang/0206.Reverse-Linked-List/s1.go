package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%d=>%s", l.Val, l.Next)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
