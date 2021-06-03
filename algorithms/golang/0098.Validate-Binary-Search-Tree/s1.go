package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	inOrder := make([]int, 0)
	getInOrder(&inOrder, root)
	// fmt.Printf("%v", inOrder)
	for i := 1; i < len(inOrder); i++ {
		if inOrder[i-1] >= inOrder[i] {
			return false
		}
	}
	return true
}

func getInOrder(inOrder *[]int, root *TreeNode) {
	if root.Left != nil {
		getInOrder(inOrder, root.Left)
	}
	*inOrder = append(*inOrder, root.Val)
	if root.Right != nil {
		getInOrder(inOrder, root.Right)
	}
}
