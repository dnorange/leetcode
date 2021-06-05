package main

import "testing"

func TestS1(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		}}

	t.Log(reverseList(list))
}
