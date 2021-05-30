package main

//func main() {
//	mini := minimumTotal2([][]int{
//		[]int{2},
//		[]int{3, 4},
//		[]int{6, 5, 7},
//		[]int{4, 1, 8, 3},
//	})
//	fmt.Printf("minimumTotal is %d\n", mini)
//}

func minimumTotal2(triangle [][]int) int {
	// check
	if len(triangle) == 0 {
		return 0
	}

	// run
	return dfs2(triangle, len(triangle)-1, 0)
}

func dfs2(triangle [][]int, i, j int) int {
	// terminator
	if i == 0 {
		return triangle[0][0]
	}

	// process
	min := 0
	if triangle[i][j] < triangle[i][j+1] {
		min = triangle[i][j]
	} else {
		min = triangle[i][j+1]
	}
	triangle[i-1][j] += min

	// clear states
	if j == len(triangle[i])-2 {
		i = i - 1
		j = -1
	}

	// drill down
	return dfs2(triangle, i, j+1)
}
