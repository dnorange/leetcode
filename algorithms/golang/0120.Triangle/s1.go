package main

import "fmt"

func main() {
	mini := minimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	})
	fmt.Printf("minimumTotal is %d\n", mini)
}

func minimumTotal(triangle [][]int) int {
	// check
	if len(triangle) == 0 {
		return 0
	}

	// run
	return dfs(triangle, 0, 0, "", 0)
}

func dfs(triangle [][]int, i, j int, path string, sum int) int {
	// process
	path += fmt.Sprintf("%d -> ", triangle[i][j])
	sum += triangle[i][j]

	// terminator
	if i == len(triangle)-1 {
		fmt.Printf("%s # sum %d\n", path, sum)
		return sum
	}

	// drill down
	leftNum := dfs(triangle, i+1, j, path, sum)
	rightNum := dfs(triangle, i+1, j+1, path, sum)

	// clear states
	// NOTE: no need clear up the state `path`
	if leftNum < rightNum {
		sum = leftNum
	} else {
		sum = rightNum
	}

	return sum
}
