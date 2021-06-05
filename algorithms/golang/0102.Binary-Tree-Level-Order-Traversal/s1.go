package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	layers := make([][]int, 0)
	put(0, root, &layers)

	return layers
}

func put(layer int, root *TreeNode, layers *[][]int) {
	if root == nil {
		return
	}
	checkLayerExist(layer, layers)
	(*layers)[layer] = append((*layers)[layer], root.Val)
	// fmt.Printf("%v", layers)
	put(layer+1, root.Left, layers)
	put(layer+1, root.Right, layers)
}

func checkLayerExist(layer int, layers *[][]int) {
	if len(*layers)-1 < layer {
		(*layers) = append((*layers), []int{})
	}
}
