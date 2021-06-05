package main

import "testing"

func TestS1(t *testing.T) {
	ret := levelOrder(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	})
	t.Log(ret)
}
