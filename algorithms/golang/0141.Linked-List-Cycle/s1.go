package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	return isMatch(head.Next, head.Next.Next)
}

func isMatch(i, j *ListNode) bool {
	if j == nil || j.Next == nil {
		return false
	}
	if i == j {
		return true
	}

	return isMatch(i.Next, j.Next.Next)
}
