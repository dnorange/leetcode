package main

import "testing"

func TestS1(t *testing.T) {
	ret := isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
	t.Log(ret)
}
